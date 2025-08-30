package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/groot314/root-g-ssh/pkg/tui"
)

const (
	host = "localhost"
	port = "3030"
)

func main() {
	model := tui.NewModel()
	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
