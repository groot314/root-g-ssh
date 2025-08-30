package tui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	contentMap map[string]string
	leftNav    table.Model
	body       string
	options    []string
	cursor     int
	message    []string
}

func NewModel() model {
	return model{
		leftNav: NavInit(),
		contentMap: map[string]string{
			"About":   "This is about me",
			"Website": "Root-G.com",
		},
	}

}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	m, _ = m.NavUpdate(msg)
	m, _ = m.ContentUpdate(msg)

	return m, nil
}

func (m model) View() string {

	header := m.HeaderView()
	body := m.ContentView()
	footer := m.FooterView()

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
		footer,
	)
}
