package cmd

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mattn/go-runewidth"
)

const (
	defaultWidth    int     = 80
	maxWidthPercent float64 = 0.8
)

type keymap struct {
	Abort key.Binding
}

var (
	_ tea.Model = (*model)(nil)
)

type model struct {
	startAt  time.Time
	duration time.Duration
	abort    bool
	progress progress.Model
	keymap   keymap
}

type modelConfig struct {
	Duration time.Duration
	Color    string
	Gradient []string
}

func newModel(cfg modelConfig) *model {
	opts := []progress.Option{progress.WithSolidFill(cfg.Color)}
	if len(cfg.Gradient) == 2 {
		opts = append(opts, progress.WithGradient(cfg.Gradient[0], cfg.Gradient[1]))
	}

	return &model{
		duration: cfg.Duration,
		abort:    false,
		progress: progress.New(opts...),
		keymap: keymap{
			Abort: key.NewBinding(key.WithKeys("ctrl+c")),
		},
	}
}

func (m *model) Init() tea.Cmd {
	runewidth.DefaultCondition.EastAsianWidth = false
	m.startAt = time.Now()
	return tea.Batch(
		m.sleep(),
		m.tick(),
	)
}

func (m *model) View() string {
	return m.progress.View()
}

type sleptMsg struct{}
type tickMsg struct{}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case sleptMsg:
		return m, tea.Quit
	case tickMsg:
		d := time.Since(m.startAt)
		cmd := m.progress.SetPercent(float64(d) / float64(m.duration))
		return m, tea.Batch(cmd, m.tick())
	case progress.FrameMsg:
		pm, cmd := m.progress.Update(msg)
		m.progress = pm.(progress.Model)
		return m, cmd
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.Abort):
			m.abort = true
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		maxWidth := int(float64(msg.Width) * maxWidthPercent)
		m.progress.Width = min(maxWidth, defaultWidth)
		return m, nil
	}

	return m, nil
}

func (m *model) sleep() tea.Cmd {
	return tea.Tick(m.duration, func(t time.Time) tea.Msg {
		return sleptMsg{}
	})
}

func (m *model) tick() tea.Cmd {
	return tea.Tick(time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
