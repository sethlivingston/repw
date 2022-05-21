package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sethlivingston/repw/internal/config"
)

type AboutView struct{}

func (v AboutView) Update(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", " ", "enter":
			m.ui.view = Repos
			return m, nil
		}
	}

	return m, nil
}

var aboutTitle lipgloss.Style = lipgloss.NewStyle().
	PaddingTop(2).
	Align(lipgloss.Center).
	Bold(true)
var aboutBody lipgloss.Style = lipgloss.NewStyle().
	Align(lipgloss.Center)

func (v AboutView) View(m Model) string {
	titleStyle := aboutTitle.Width(m.ui.termWidth)
	bodyStyle := aboutBody.Width(m.ui.termWidth)

	s := fmt.Sprintln(titleStyle.Render("repw"))
	s += fmt.Sprintln(bodyStyle.Render(config.Version))

	return s
}
