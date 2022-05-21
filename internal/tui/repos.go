package tui

import tea "github.com/charmbracelet/bubbletea"

type ReposView struct{}

func (v ReposView) Update(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "a":
			m.ui.view = About
			return m, nil
		case ",":
			m.ui.view = Settings
			return m, nil
		}
	}

	return m, nil
}

func (v ReposView) View(m Model) string {
	s := "Repos view"

	return s
}
