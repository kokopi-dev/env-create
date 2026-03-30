package tui

import (
	"env-create/internal/pages"
	"env-create/internal/services"
	"os"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

func (m ScpTUIInterface) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case pages.ScpCheckCacheMsg:
		if m.Services.Config == nil {
			return m, func() tea.Msg { return pages.ScpPageMsg{} }
		}
		svc := m.Services.Config
		return m, func() tea.Msg {
			if !svc.CacheExists() {
				return pages.ScpCacheCheckedMsg{CacheExists: false}
			}
			cwd, _ := os.Getwd()
			cfg, err := svc.LoadConfig(cwd)
			var prefill *pages.ScpPrefill
			if err == nil && cfg != nil {
				prefill = &pages.ScpPrefill{
					Username:   cfg.Username,
					Host:       cfg.Host,
					RemotePath: cfg.RemotePath,
				}
			}
			return pages.ScpCacheCheckedMsg{CacheExists: true, Prefill: prefill}
		}

	case pages.ScpCacheCheckedMsg:
		if !msg.CacheExists {
			m.Page = scpPageCachePrompt
			return m, nil
		}
		return m, func() tea.Msg {
			return pages.ScpPageMsg{Prefill: msg.Prefill, SaveConfigs: true}
		}

	case pages.ScpPageMsg:
		m.Page = scpPageForm
		m.SaveConfigs = msg.SaveConfigs
		m.Inputs[fieldUsername] = pages.InitUsernameInput()
		m.Inputs[fieldHost] = pages.InitHostInput()
		m.Inputs[fieldRemotePath] = pages.InitRemotePathInput()
		if msg.Prefill != nil {
			m.Inputs[fieldUsername].SetValue(msg.Prefill.Username)
			m.Inputs[fieldHost].SetValue(msg.Prefill.Host)
			m.Inputs[fieldRemotePath].SetValue(msg.Prefill.RemotePath)
		}
		m.Inputs[fieldUsername].Focus()
		return m, textinput.Blink

	case pages.ScpResultMsg:
		m.Page = scpPageResult
		m.ResultOutput = msg.Output
		m.ResultErr = msg.Err
		if m.SaveConfigs && m.Services.Config != nil {
			u := m.Inputs[fieldUsername].Value()
			h := m.Inputs[fieldHost].Value()
			p := m.Inputs[fieldRemotePath].Value()
			return m, func() tea.Msg {
				return pages.ScpSaveConfigMsg{Username: u, Host: h, RemotePath: p}
			}
		}
		return m, nil

	case pages.ScpSaveConfigMsg:
		if m.Services.Config != nil {
			cwd, _ := os.Getwd()
			_ = m.Services.Config.SaveConfig(cwd, services.ScpConfig{
				Username:   msg.Username,
				Host:       msg.Host,
				RemotePath: msg.RemotePath,
			})
		}
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

		case "y", "Y":
			if m.Page == scpPageCachePrompt {
				svc := m.Services.Config
				return m, func() tea.Msg {
					if err := svc.CreateCacheDir(); err != nil {
						return pages.ScpPageMsg{SaveConfigs: false}
					}
					return pages.ScpPageMsg{SaveConfigs: true}
				}
			}

		case "n", "N":
			if m.Page == scpPageCachePrompt {
				return m, func() tea.Msg { return pages.ScpPageMsg{SaveConfigs: false} }
			}

		case "enter":
			if m.Page == scpPageResult {
				m.Accepted = true
				m.Quitting = true
				return m, tea.Quit
			}
			if m.Page == scpPageCachePrompt {
				return m, nil
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

	if m.Page != scpPageForm {
		return m, nil
	}
	var cmd tea.Cmd
	m.Inputs[m.ActiveField], cmd = m.Inputs[m.ActiveField].Update(msg)
	return m, cmd
}
