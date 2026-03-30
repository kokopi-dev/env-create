package tui

import (
	"env-create/internal/pages"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

func (m ScpTUIInterface) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case pages.ScpPageMsg:
		m.Inputs[fieldUsername] = pages.InitUsernameInput()
		m.Inputs[fieldHost] = pages.InitHostInput()
		m.Inputs[fieldRemotePath] = pages.InitRemotePathInput()
		m.Inputs[fieldUsername].Focus()
		return m, textinput.Blink

	case pages.ScpResultMsg:
		m.ShowResult = true
		m.ResultOutput = msg.Output
		m.ResultErr = msg.Err
		return m, nil

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
			if m.ShowResult {
				m.Accepted = true
				m.Quitting = true
				return m, tea.Quit
			}
			if m.ActiveField < fieldCount-1 {
				m.Inputs[m.ActiveField].Blur()
				m.ActiveField++
				m.Inputs[m.ActiveField].Focus()
				return m, textinput.Blink
			}
			// last field — run scp asynchronously
			username := m.Inputs[fieldUsername].Value()
			host := m.Inputs[fieldHost].Value()
			path := m.Inputs[fieldRemotePath].Value()
			svc := m.Services.Scp
			return m, func() tea.Msg {
				out, err := svc.Run(username, host, path)
				return pages.ScpResultMsg{Output: out, Err: err}
			}
		}
	}

	var cmd tea.Cmd
	m.Inputs[m.ActiveField], cmd = m.Inputs[m.ActiveField].Update(msg)
	return m, cmd
}
