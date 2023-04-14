package cmd

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
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

func newModel(t float64) *model {
	return &model{
		duration: time.Duration(t*1000) * time.Millisecond,
		abort:    false,
		progress: progress.New(progress.WithDefaultGradient()),
		keymap: keymap{
			Abort: key.NewBinding(key.WithKeys("ctrl+c")),
		},
	}
}

func (m *model) Init() tea.Cmd {
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
