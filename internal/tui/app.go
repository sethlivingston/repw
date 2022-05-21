package tui

import tea "github.com/charmbracelet/bubbletea"

type View interface {
	Update(m Model, msg tea.Msg) (tea.Model, tea.Cmd)
	View(m Model) string
}

type ViewName string

const (
	Repos ViewName = "repos"
	About ViewName = "about"
)

var AppViews = map[ViewName]View{
	Repos: ReposView{},
	About: AboutView{},
}

type ModelUI struct {
	view       ViewName
	termWidth  int
	termHeight int
}

type Model struct {
	ui ModelUI
}

func NewModel() Model {
	return Model{
		ui: ModelUI{
			view: Repos,
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.ui.termWidth = msg.Width
		m.ui.termHeight = msg.Height
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return AppViews[m.ui.view].Update(m, msg)
}

func (m Model) View() string {
	return AppViews[m.ui.view].View(m)
}
