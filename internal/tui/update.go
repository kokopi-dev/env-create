package tui

import (
	"env-create/internal/pages"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

func (m TUIInterface) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case pages.ProjectNamePageMsg:
		m.Input = pages.InitProjectNamePage(m.Services)
		return m, textinput.Blink

	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
			m.Accepted = true
			m.Quitting = true
			return m, tea.Quit
		case "ctrl+c", "esc":
			m.Quitting = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.Input, cmd = m.Input.Update(msg)
	return m, cmd
}
