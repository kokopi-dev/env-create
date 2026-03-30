package tui

import (
	"env-create/internal/pages"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

// main page render loop/flow
func (m TUIInterface) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// pages
	case pages.HomePageMsg:
		m.Pages.Home = pages.InitHomePage()
		m.Page = "home"
		return m, nil

	case pages.ProjectNamePageMsg:
		m.Pages.ProjectName = pages.InitProjectNamePage(m.Services)
		m.Page = "project-name"
		return m, textinput.Blink

	case pages.ProjectNameAcceptedMsg:
		m.Accepted = true
		m.AcceptedValue = msg.Value
		m.Quitting = true
		return m, tea.Quit

	case pages.SendSelectedMsg:
		m.Accepted = true
		m.AcceptedValue = "send"
		m.Quitting = true
		return m, tea.Quit

	// general
	case tea.WindowSizeMsg:
		m.WindowWidth = msg.Width
		m.WindowHeight = msg.Height
		return m, nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.Quitting = true
			return m, tea.Quit
		}
	}

	// user output
	var cmd tea.Cmd
	switch m.Page {
	case "home":
		m.Pages.Home, cmd = m.Pages.Home.Update(msg)
	case "project-name":
		m.Pages.ProjectName, cmd = m.Pages.ProjectName.Update(msg)
	}
	return m, cmd
}
