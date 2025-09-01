package tui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	contentMap map[string]string
	leftNav    table.Model
	body       viewport.Model
}

func NewModel() model {
	return model{
		leftNav: NavInit(),
		contentMap: map[string]string{
			"About":   "This is about me",
			"Website": "Root-G.com",
			"Projects": "Current project list:\n" +
				"- About Me CLI via SSH\n" +
				"- Website: Root-G.com\n" +
				"- Homelab\n" +
				"- NixOS\n",
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

var borderStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func (m model) View() string {

	header := m.HeaderView()
	nav := m.NavView()
	body := m.ContentView()
	footer := m.FooterView()

	main := lipgloss.JoinHorizontal(lipgloss.Top, nav, body)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		main,
		footer,
	)
}
