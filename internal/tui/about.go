package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sethlivingston/repw/internal/config"
	"github.com/sethlivingston/repw/internal/tui/styles"
)

type AboutView struct{}

func (v AboutView) Update(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.ui.view = Repos
			return m, nil
		}
	}

	return m, nil
}

func (v AboutView) View(m Model) string {
	titleStyle := styles.Title.Width(m.ui.termWidth)
	subtitleStyle := styles.Subtitle.Width(m.ui.termWidth)

	s := fmt.Sprintln(titleStyle.Render("repw"))
	s += fmt.Sprintln(subtitleStyle.Render(config.Version))

	return s
}
