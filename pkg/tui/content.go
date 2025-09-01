package tui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) ContentView() string {
	var contentStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		Align(lipgloss.Left)

	return contentStyle.Render(m.body.View())
}

func (m model) ContentUpdate(msg tea.Msg) (model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:

		navWidth := lipgloss.Width(m.NavView())

		verticalMarginHeight := lipgloss.Height(m.HeaderView()) + lipgloss.Height(m.FooterView())

		m.body = viewport.New(msg.Width-navWidth-2, min(msg.Height-verticalMarginHeight-2, 40))
		if b, ok := m.contentMap[m.leftNav.SelectedRow()[0]]; ok {
			m.body.SetContent(b)
		}
	}

	if b, ok := m.contentMap[m.leftNav.SelectedRow()[0]]; ok {
		m.body.SetContent(b)
	}

	if !m.leftNav.Focused() {
		m.body, _ = m.body.Update(msg)
	}

	return m, nil
}
