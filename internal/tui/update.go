package tui

import (
	"env-create/internal/pages"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

func (m TUIInterface) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case pages.HomePageMsg:
		m.Home = pages.InitHomePage()
		m.Page = "home"
		return m, nil

	case pages.ProjectNamePageMsg:
		m.Input = pages.InitProjectNamePage(m.Services)
		m.Page = "project-name"
		return m, textinput.Blink

	case tea.WindowSizeMsg:
		m.WindowWidth = msg.Width
		m.WindowHeight = msg.Height
		return m, nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.Quitting = true
			return m, tea.Quit
		case "enter":
			if m.Page == "project-name" {
				m.Accepted = true
				m.Quitting = true
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	if m.Page == "home" {
		m.Home, cmd = m.Home.Update(msg)
	} else {
		m.Input, cmd = m.Input.Update(msg)
	}
	return m, cmd
}
