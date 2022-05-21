package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Key bindings

type appKeyMap struct {
	Help key.Binding
	Quit key.Binding
}

func (k appKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k appKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{},
		{k.Help, k.Quit},
	}
}

var keys = appKeyMap{
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "togglehelp"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}

// Views

type View interface {
	Update(m Model, msg tea.Msg) (tea.Model, tea.Cmd)
	View(m Model) string
}

type ViewName string

const (
	Repos    ViewName = "repos"
	About    ViewName = "about"
	Settings ViewName = "settings"
)

var AppViews = map[ViewName]View{
	Repos:    ReposView{},
	About:    AboutView{},
	Settings: SettingsView{},
}

// Model

type ModelUI struct {
	View       ViewName
	TermWidth  int
	TermHeight int
	Keys       appKeyMap
	Help       help.Model
}

type Model struct {
	UI ModelUI
}

func NewModel() Model {
	return Model{
		UI: ModelUI{
			View: Repos,
			Keys: keys,
			Help: help.New(),
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Application update

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.UI.TermWidth = msg.Width
		m.UI.TermHeight = msg.Height
		m.UI.Help.Width = msg.Width
		return m, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.UI.Keys.Help):
			m.UI.Help.ShowAll = !m.UI.Help.ShowAll
			return m, nil
		case key.Matches(msg, m.UI.Keys.Quit):
			return m, tea.Quit
		}
	}

	return AppViews[m.UI.View].Update(m, msg)
}

// Application view

func (m Model) View() string {
	appView := AppViews[m.UI.View].View(m)
	appViewHeight := strings.Count(appView, "\n")
	helpView := m.UI.Help.View(m.UI.Keys)

	if m.UI.TermHeight > 0 {
		return appView + lipgloss.PlaceVertical(m.UI.TermHeight-appViewHeight, lipgloss.Bottom, helpView)
	} else {
		return appView + "\n\n" + helpView
	}

}
