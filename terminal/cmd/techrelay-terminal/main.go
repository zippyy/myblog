package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/wish/v2"
	"charm.land/wish/v2/activeterm"
	"charm.land/wish/v2/bubbletea"
	"charm.land/wish/v2/logging"
	"github.com/charmbracelet/ssh"
)

const (
	defaultFeedURL     = "https://techrelay.xyz/index.json"
	defaultSiteURL     = "https://techrelay.xyz"
	defaultListenAddr  = "0.0.0.0:23234"
	defaultHealthAddr  = "0.0.0.0:8080"
	defaultHostKeyPath = "/data/ssh_host_ed25519"
)

func main() {
	feedURL := env("TECHRELAY_FEED_URL", defaultFeedURL)
	siteURL := strings.TrimRight(env("TECHRELAY_SITE_URL", defaultSiteURL), "/")
	listenAddr := env("TECHRELAY_LISTEN_ADDR", defaultListenAddr)
	healthAddr := env("TECHRELAY_HEALTH_ADDR", defaultHealthAddr)
	hostKeyPath := env("TECHRELAY_HOST_KEY", defaultHostKeyPath)
	refreshEvery := envDuration("TECHRELAY_REFRESH_INTERVAL", 15*time.Minute)

	contentCache := &cache{}
	posts, err := fetchPosts(feedURL)
	contentCache.replace(posts, err)
	if err != nil {
		log.Printf("initial feed fetch failed: %v", err)
	} else {
		log.Printf("loaded %d posts from %s", len(posts), feedURL)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go refreshLoop(ctx, contentCache, feedURL, refreshEvery)
	go serveHealth(ctx, healthAddr, contentCache)

	server, err := wish.NewServer(
		wish.WithAddress(listenAddr),
		wish.WithHostKeyPath(hostKeyPath),
		wish.WithMiddleware(
			bubbletea.Middleware(func(session ssh.Session) (tea.Model, []tea.ProgramOption) {
				pty, _, _ := session.Pty()
				items, updatedAt, cacheErr := contentCache.snapshot()
				m := newModel(contentCache, feedURL, siteURL, items, updatedAt, cacheErr)
				m.width = pty.Window.Width
				m.height = pty.Window.Height
				return m, nil
			}),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Fatalf("create SSH server: %v", err)
	}

	go func() {
		log.Printf("Tech Relay terminal listening on %s", listenAddr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Printf("SSH server failed: %v", err)
			cancel()
		}
	}()

	<-ctx.Done()
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := server.Shutdown(shutdownCtx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Printf("SSH shutdown failed: %v", err)
	}
}

func env(name, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(name)); value != "" {
		return value
	}
	return fallback
}

func envDuration(name string, fallback time.Duration) time.Duration {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return fallback
	}
	parsed, err := time.ParseDuration(value)
	if err != nil {
		log.Printf("invalid %s=%q, using %s", name, value, fallback)
		return fallback
	}
	return parsed
}
