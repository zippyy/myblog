package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	tea "charm.land/bubbletea/v2"
)

type post struct {
	Title       string   `json:"title"`
	Date        string   `json:"date"`
	LastMod     string   `json:"lastmod"`
	URL         string   `json:"url"`
	Path        string   `json:"path"`
	Description string   `json:"description"`
	Summary     string   `json:"summary"`
	Content     string   `json:"content"`
	Categories  []string `json:"categories"`
	Tags        []string `json:"tags"`
	Featured    bool     `json:"featured"`
	ReadingTime int      `json:"readingTime"`
	WordCount   int      `json:"wordCount"`
}

type cache struct {
	mu        sync.RWMutex
	posts     []post
	updatedAt time.Time
	lastErr   error
}

func (c *cache) snapshot() ([]post, time.Time, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	items := make([]post, len(c.posts))
	copy(items, c.posts)
	return items, c.updatedAt, c.lastErr
}

func (c *cache) replace(posts []post, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if err == nil {
		c.posts = posts
		c.updatedAt = time.Now()
	}
	c.lastErr = err
}

type fetchResultMsg struct {
	posts []post
	err   error
}

func fetchCmd(url string) tea.Cmd {
	return func() tea.Msg {
		posts, err := fetchPosts(url)
		return fetchResultMsg{posts: posts, err: err}
	}
}

func fetchPosts(url string) ([]post, error) {
	client := &http.Client{Timeout: 15 * time.Second}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "techrelay-terminal/1.0")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("feed returned %s", resp.Status)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 16<<20))
	if err != nil {
		return nil, err
	}

	var posts []post
	if err := json.Unmarshal(body, &posts); err != nil {
		var envelope struct {
			Posts []post `json:"posts"`
		}
		if envelopeErr := json.Unmarshal(body, &envelope); envelopeErr != nil {
			return nil, fmt.Errorf("decode feed: %w", err)
		}
		posts = envelope.Posts
	}
	for i := range posts {
		posts[i].Title = strings.TrimSpace(posts[i].Title)
		posts[i].Content = strings.TrimSpace(posts[i].Content)
		if posts[i].URL == "" && posts[i].Path != "" {
			posts[i].URL = strings.TrimRight(defaultSiteURL, "/") + "/" + strings.TrimLeft(posts[i].Path, "/")
		}
	}
	sort.SliceStable(posts, func(i, j int) bool { return posts[i].Date > posts[j].Date })
	return posts, nil
}

func refreshLoop(ctx context.Context, c *cache, url string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			posts, err := fetchPosts(url)
			c.replace(posts, err)
			if err != nil {
				log.Printf("background feed refresh failed: %v", err)
			} else {
				log.Printf("background feed refresh loaded %d posts", len(posts))
			}
		}
	}
}

func serveHealth(ctx context.Context, addr string, c *cache) {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		posts, updatedAt, lastErr := c.snapshot()
		w.Header().Set("Content-Type", "application/json")
		status := http.StatusOK
		if len(posts) == 0 && lastErr != nil {
			status = http.StatusServiceUnavailable
		}
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"ok":         status == http.StatusOK,
			"posts":      len(posts),
			"updated_at": updatedAt,
			"error":      errorString(lastErr),
		})
	})
	server := &http.Server{Addr: addr, Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = server.Shutdown(shutdownCtx)
	}()
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("health server failed: %v", err)
	}
}

func errorString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
