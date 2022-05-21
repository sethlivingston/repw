package styles

import "github.com/charmbracelet/lipgloss"

var Title lipgloss.Style = lipgloss.NewStyle().
	PaddingTop(1).
	Align(lipgloss.Center).
	Bold(true)

var Subtitle lipgloss.Style = lipgloss.NewStyle().
	Align(lipgloss.Center)

var Section lipgloss.Style = lipgloss.NewStyle().
	PaddingTop(1)
