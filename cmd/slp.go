package cmd

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keymap struct {
	Abort key.Binding
}

var (
	_ tea.Model = (*model)(nil)
)

type model struct {
	duration time.Duration
	abort    bool
	keymap   keymap
}

func newModel(t float64) *model {
	return &model{
		duration: time.Duration(t*1000) * time.Millisecond,
		abort:    false,
		keymap: keymap{
			Abort: key.NewBinding(key.WithKeys("ctrl+c")),
		},
	}
}

func (m *model) Init() tea.Cmd {
	return m.sleep()
}

func (m *model) View() string {
	return "slp"
}

type sleptMsg struct{}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case sleptMsg:
		return m, tea.Quit
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
