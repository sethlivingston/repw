package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sethlivingston/repw/internal/tui"
)

func main() {
	program := tea.NewProgram(tui.NewModel())
	err := program.Start()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
