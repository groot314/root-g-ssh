package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func (m model) ContentView() string {
	var contentStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).Height(20).Width(50)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		baseStyle.Render(m.NavView()),
		contentStyle.Render(m.body),
	)
}

func (m model) ContentUpdate(msg tea.Msg) (model, tea.Cmd) {
	if body, ok := m.contentMap[m.leftNav.SelectedRow()[0]]; ok {
		m.body = body
	}

	return m, nil
}
