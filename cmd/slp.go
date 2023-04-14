package cmd

import (
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
	abort  bool
	keymap keymap
}

func newModel(d float64) *model {
	return &model{
		abort: false,
		keymap: keymap{
			Abort: key.NewBinding(key.WithKeys("ctrl+c")),
		},
	}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) View() string {
	return "slp"
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.Abort):
			m.abort = true
			return m, tea.Quit
		}
	}

	return m, nil
}
