package tui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) NavView() string {
	return borderStyle.Render(m.leftNav.View())
}

func (m model) NavUpdate(msg tea.Msg) (model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if !m.leftNav.Focused() {
				m.leftNav.Focus()
			}
		case "enter":
			if m.leftNav.Focused() {
				m.leftNav.Blur()
			}
		}
	case tea.WindowSizeMsg:
		verticalMarginHeight := lipgloss.Height(m.HeaderView()) + lipgloss.Height(m.FooterView())
		m.leftNav.SetHeight(min(40, msg.Height-verticalMarginHeight-2))

	}
	m.leftNav, _ = m.leftNav.Update(msg)

	return m, nil
}

func NavInit() table.Model {
	columns := []table.Column{
		{Title: "Nav", Width: 10},
	}

	rows := []table.Row{
		{"About"},
		{"Website"},
		{"Projects"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("21")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("15")).
		Background(lipgloss.Color("21")).
		Bold(true)
	t.SetStyles(s)
	t.SetHeight(20)

	return t
}
