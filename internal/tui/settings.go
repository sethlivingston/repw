package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sethlivingston/repw/internal/tui/styles"
)

type SettingsView struct{}

func (v SettingsView) Update(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (v SettingsView) View(m Model) string {
	titleStyle := styles.Title.Width(m.ui.termWidth)

	s := fmt.Sprintln(titleStyle.Render("Settings"))

	return s

}
