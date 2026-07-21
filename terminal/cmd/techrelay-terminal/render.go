package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	brandStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#ff5fa2"))
	mutedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#7f8c8d"))
	dateStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#65d1ff"))
	activeStyle = lipgloss.NewStyle().Bold(true).
		Foreground(lipgloss.Color("#101014")).
		Background(lipgloss.Color("#ff5fa2"))
)

func (m model) View() tea.View {
	var body string
	switch m.current {
	case screenArticle:
		body = m.renderArticle()
	case screenAbout:
		body = m.renderAbout()
	case screenHelp:
		body = m.renderHelp()
	default:
		body = m.renderPosts()
	}
	view := tea.NewView(body)
	view.AltScreen = true
	return view
}

func (m model) renderPosts() string {
	width := max(36, m.width)
	var b strings.Builder
	b.WriteString(renderHeader(width))
	b.WriteByte('\n')

	if m.searching {
		b.WriteString(brandStyle.Render("Search: ") + m.query + "█\n")
	} else if m.query != "" {
		b.WriteString(fmt.Sprintf("Filter: %q  (%d matches)\n", m.query, len(m.filtered)))
	} else {
		b.WriteString(fmt.Sprintf("Latest posts  ·  %d articles\n", len(m.posts)))
	}
	b.WriteString(rule(width))
	b.WriteByte('\n')

	if len(m.filtered) == 0 {
		b.WriteString("\nNo posts match the current search. Press Esc or / to change it.\n")
	} else {
		page := m.listPageSize()
		end := min(len(m.filtered), m.listOffset+page)
		for visibleIndex := m.listOffset; visibleIndex < end; visibleIndex++ {
			item := m.posts[m.filtered[visibleIndex]]
			date := displayDate(item.Date)
			titleWidth := max(12, width-16)
			title := truncate(item.Title, titleWidth)
			line := fmt.Sprintf("%-10s  %s", date, title)
			if visibleIndex == m.cursor {
				b.WriteString(activeStyle.Render(" " + line + " "))
			} else {
				b.WriteString(dateStyle.Render(date) + "  " + title)
			}
			b.WriteByte('\n')
			summary := firstNonEmpty(item.Description, item.Summary)
			b.WriteString("  " + mutedStyle.Render(truncate(cleanSpace(summary), max(12, width-4))))
			b.WriteByte('\n')
			b.WriteByte('\n')
		}
	}

	footer := "↑/↓ or j/k navigate  enter read  / search  a about  ? help  r refresh  q quit"
	if m.loading {
		footer = "Refreshing…  " + footer
	} else if m.status != "" {
		footer = m.status + "  ·  " + footer
	}
	b.WriteString(rule(width))
	b.WriteByte('\n')
	b.WriteString(mutedStyle.Render(truncate(footer, width)))
	return fitHeight(b.String(), m.height)
}

func (m model) renderArticle() string {
	item, ok := m.selectedPost()
	if !ok {
		return m.renderPosts()
	}
	width := max(36, m.width)
	contentWidth := max(24, width-4)
	var b strings.Builder
	b.WriteString(brandStyle.Render(truncate(item.Title, width)))
	b.WriteByte('\n')
	meta := displayDate(item.Date)
	if item.ReadingTime > 0 {
		meta += fmt.Sprintf("  ·  %d min read", item.ReadingTime)
	}
	if len(item.Categories) > 0 {
		meta += "  ·  " + strings.Join(item.Categories, ", ")
	}
	b.WriteString(mutedStyle.Render(truncate(meta, width)))
	b.WriteByte('\n')
	b.WriteString(dateStyle.Render(truncate(item.URL, width)))
	b.WriteByte('\n')
	b.WriteString(rule(width))
	b.WriteByte('\n')

	lines := articleLines(item, contentWidth)
	pageSize := m.articlePageSize()
	end := min(len(lines), m.articleOffset+pageSize)
	for _, line := range lines[m.articleOffset:end] {
		b.WriteString(line)
		b.WriteByte('\n')
	}
	for i := end - m.articleOffset; i < pageSize; i++ {
		b.WriteByte('\n')
	}

	percent := 100
	if len(lines) > pageSize {
		percent = min(100, int(float64(m.articleOffset+pageSize)/float64(len(lines))*100))
	}
	b.WriteString(rule(width))
	b.WriteByte('\n')
	b.WriteString(mutedStyle.Render(truncate(fmt.Sprintf("j/k scroll  space/PgDn page  g/G top/bottom  b back  ctrl+c quit  ·  %d%%", percent), width)))
	return fitHeight(b.String(), m.height)
}

func (m model) renderAbout() string {
	width := max(36, m.width)
	content := []string{
		brandStyle.Render("ABOUT TECH RELAY"),
		"",
		"Tech Relay covers practical IT, networking, automation, self-hosting,",
		"remote work, and the tools that make technology less painful.",
		"",
		"Website: " + m.siteURL,
		"Terminal feed: " + m.feedURL,
		"",
		"This SSH interface reads the same Hugo content used by the website.",
		"Publishing a new post updates both versions automatically.",
		"",
		mutedStyle.Render("Press b, Esc, Enter, or q to return."),
	}
	return fitHeight(wrapBlock(strings.Join(content, "\n"), width), m.height)
}

func (m model) renderHelp() string {
	width := max(36, m.width)
	content := []string{
		brandStyle.Render("KEYBOARD HELP"),
		"",
		"Post list",
		"  ↑/↓ or j/k       move selection",
		"  PgUp/PgDn        move one page",
		"  Enter            read selected post",
		"  /                search titles, summaries, categories, and tags",
		"  r                reload the Hugo JSON feed",
		"  a                about Tech Relay",
		"  q                disconnect",
		"",
		"Article reader",
		"  ↑/↓ or j/k       scroll one line",
		"  Space/PgDn       scroll one page",
		"  g / G            jump to top / bottom",
		"  b, Esc, or q     return to the post list",
		"  Ctrl+C           disconnect immediately",
		"",
		mutedStyle.Render("Press b, Esc, Enter, or q to return."),
	}
	return fitHeight(wrapBlock(strings.Join(content, "\n"), width), m.height)
}

func articleLines(item post, width int) []string {
	content := strings.TrimSpace(item.Content)
	if content == "" {
		content = firstNonEmpty(item.Description, item.Summary, "This post has no terminal-readable body yet.")
	}
	paragraphs := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")
	lines := make([]string, 0, len(paragraphs)*2)
	for _, paragraph := range paragraphs {
		paragraph = cleanSpace(paragraph)
		if paragraph == "" {
			if len(lines) == 0 || lines[len(lines)-1] != "" {
				lines = append(lines, "")
			}
			continue
		}
		lines = append(lines, wrapWords(paragraph, width)...)
		lines = append(lines, "")
	}
	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	if len(lines) == 0 {
		return []string{"No content available."}
	}
	return lines
}

func renderHeader(width int) string {
	name := " TECH RELAY "
	tagline := "IT · NETWORKING · AUTOMATION · SELF-HOSTING"
	if width < 58 {
		tagline = "IT · NETWORKING · SELF-HOSTING"
	}
	return brandStyle.Render(center(name, width)) + "\n" + mutedStyle.Render(center(tagline, width))
}

func rule(width int) string {
	return mutedStyle.Render(strings.Repeat("─", max(1, width)))
}

func displayDate(value string) string {
	if len(value) >= 10 {
		return value[:10]
	}
	return value
}

func wrapBlock(value string, width int) string {
	var out []string
	for _, line := range strings.Split(value, "\n") {
		if strings.TrimSpace(line) == "" {
			out = append(out, "")
			continue
		}
		out = append(out, wrapWords(line, width)...)
	}
	return strings.Join(out, "\n")
}

func wrapWords(value string, width int) []string {
	width = max(1, width)
	words := strings.Fields(value)
	if len(words) == 0 {
		return []string{""}
	}
	lines := []string{}
	line := words[0]
	for _, word := range words[1:] {
		if runeLen(line)+1+runeLen(word) <= width {
			line += " " + word
			continue
		}
		lines = append(lines, line)
		for runeLen(word) > width {
			runes := []rune(word)
			lines = append(lines, string(runes[:width]))
			word = string(runes[width:])
		}
		line = word
	}
	return append(lines, line)
}

func truncate(value string, width int) string {
	value = cleanSpace(value)
	if width <= 0 {
		return ""
	}
	runes := []rune(value)
	if len(runes) <= width {
		return value
	}
	if width == 1 {
		return "…"
	}
	return string(runes[:width-1]) + "…"
}

func fitHeight(value string, height int) string {
	if height <= 0 {
		return value
	}
	lines := strings.Split(value, "\n")
	if len(lines) > height {
		lines = lines[:height]
	}
	return strings.Join(lines, "\n")
}

func center(value string, width int) string {
	length := runeLen(value)
	if length >= width {
		return truncate(value, width)
	}
	left := (width - length) / 2
	return strings.Repeat(" ", left) + value
}

func cleanSpace(value string) string {
	return strings.Join(strings.Fields(value), " ")
}

func runeLen(value string) int {
	return utf8.RuneCountInString(value)
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
