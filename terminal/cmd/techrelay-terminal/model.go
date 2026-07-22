package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	tea "charm.land/bubbletea/v2"
)

type screen int

const (
	screenPosts screen = iota
	screenArticle
	screenAbout
	screenHelp
)

type model struct {
	cache         *cache
	feedURL       string
	siteURL       string
	posts         []post
	filtered      []int
	current       screen
	cursor        int
	listOffset    int
	articleOffset int
	query         string
	searching     bool
	loading       bool
	status        string
	width         int
	height        int
	updatedAt     time.Time
}

func newModel(c *cache, feedURL, siteURL string, posts []post, updatedAt time.Time, cacheErr error) model {
	m := model{
		cache:     c,
		feedURL:   feedURL,
		siteURL:   siteURL,
		posts:     posts,
		current:   screenPosts,
		width:     80,
		height:    24,
		updatedAt: updatedAt,
	}
	if cacheErr != nil {
		m.status = "Feed error: " + cacheErr.Error()
	}
	m.applyFilter()
	return m
}

func (m model) Init() tea.Cmd {
	if len(m.posts) == 0 {
		return tea.Batch(tea.RequestBackgroundColor, fetchCmd(m.feedURL))
	}
	return tea.RequestBackgroundColor
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.clamp()
		return m, nil
	case fetchResultMsg:
		m.loading = false
		if msg.err != nil {
			m.status = "Refresh failed: " + msg.err.Error()
			m.cache.replace(nil, msg.err)
			return m, nil
		}
		m.posts = msg.posts
		m.updatedAt = time.Now()
		m.status = fmt.Sprintf("Loaded %d posts", len(msg.posts))
		m.cache.replace(msg.posts, nil)
		m.applyFilter()
		return m, nil
	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if m.searching {
			return m.handleSearch(msg)
		}
		return m.handleKey(msg)
	}
	return m, nil
}

func (m model) handleSearch(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.searching = false
		m.query = ""
		m.applyFilter()
	case "enter":
		m.searching = false
	case "backspace":
		if len(m.query) > 0 {
			_, size := utf8.DecodeLastRuneInString(m.query)
			m.query = m.query[:len(m.query)-size]
			m.applyFilter()
		}
	case "ctrl+w":
		m.query = strings.TrimRight(m.query, " ")
		if i := strings.LastIndexByte(m.query, ' '); i >= 0 {
			m.query = m.query[:i+1]
		} else {
			m.query = ""
		}
		m.applyFilter()
	default:
		if text := msg.Key().Text; text != "" {
			m.query += text
			m.applyFilter()
		}
	}
	return m, nil
}

func (m model) handleKey(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	key := msg.String()
	switch m.current {
	case screenPosts:
		switch key {
		case "q":
			return m, tea.Quit
		case "j", "down":
			m.cursor++
		case "k", "up":
			m.cursor--
		case "pgdown", "ctrl+f":
			m.cursor += m.listPageSize()
		case "pgup", "ctrl+b":
			m.cursor -= m.listPageSize()
		case "home", "g":
			m.cursor = 0
		case "end", "G":
			m.cursor = len(m.filtered) - 1
		case "enter":
			if len(m.filtered) > 0 {
				m.current = screenArticle
				m.articleOffset = 0
			}
		case "/":
			m.searching = true
		case "a":
			m.current = screenAbout
		case "?":
			m.current = screenHelp
		case "r":
			m.loading = true
			m.status = "Refreshing feed…"
			return m, fetchCmd(m.feedURL)
		}
	case screenArticle:
		switch key {
		case "q", "b", "esc":
			m.current = screenPosts
		case "j", "down":
			m.articleOffset++
		case "k", "up":
			m.articleOffset--
		case "pgdown", "ctrl+f", " ":
			m.articleOffset += m.articlePageSize()
		case "pgup", "ctrl+b":
			m.articleOffset -= m.articlePageSize()
		case "home", "g":
			m.articleOffset = 0
		case "end", "G":
			m.articleOffset = m.maxArticleOffset()
		}
	case screenAbout, screenHelp:
		switch key {
		case "q", "b", "esc", "enter":
			m.current = screenPosts
		}
	}
	m.clamp()
	return m, nil
}

func (m *model) applyFilter() {
	needle := strings.ToLower(strings.TrimSpace(m.query))
	m.filtered = m.filtered[:0]
	for i, item := range m.posts {
		haystack := strings.ToLower(strings.Join([]string{
			item.Title,
			item.Description,
			item.Summary,
			strings.Join(item.Categories, " "),
			strings.Join(item.Tags, " "),
		}, " "))
		if needle == "" || strings.Contains(haystack, needle) {
			m.filtered = append(m.filtered, i)
		}
	}
	m.cursor = 0
	m.listOffset = 0
	m.clamp()
}

func (m *model) clamp() {
	if len(m.filtered) == 0 {
		m.cursor = 0
		m.listOffset = 0
	} else {
		if m.cursor < 0 {
			m.cursor = 0
		}
		if m.cursor >= len(m.filtered) {
			m.cursor = len(m.filtered) - 1
		}
		page := m.listPageSize()
		if m.cursor < m.listOffset {
			m.listOffset = m.cursor
		}
		if m.cursor >= m.listOffset+page {
			m.listOffset = m.cursor - page + 1
		}
		maxOffset := max(0, len(m.filtered)-page)
		if m.listOffset > maxOffset {
			m.listOffset = maxOffset
		}
	}

	if m.articleOffset < 0 {
		m.articleOffset = 0
	}
	if maxOffset := m.maxArticleOffset(); m.articleOffset > maxOffset {
		m.articleOffset = maxOffset
	}
}

func (m model) selectedPost() (post, bool) {
	if len(m.filtered) == 0 || m.cursor < 0 || m.cursor >= len(m.filtered) {
		return post{}, false
	}
	return m.posts[m.filtered[m.cursor]], true
}

func (m model) listPageSize() int {
	return max(1, (max(12, m.height)-8)/3)
}

func (m model) articlePageSize() int {
	return max(3, max(12, m.height)-6)
}

func (m model) maxArticleOffset() int {
	item, ok := m.selectedPost()
	if !ok {
		return 0
	}
	lines := articleLines(item, max(24, max(36, m.width)-4))
	return max(0, len(lines)-m.articlePageSize())
}
