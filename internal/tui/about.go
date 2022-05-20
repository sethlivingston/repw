package tui

import tea "github.com/charmbracelet/bubbletea"

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

func (v AboutView) View(m Model) string {
	s := "About repw!"

	return s
}
