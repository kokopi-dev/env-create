package tui

import (
	"env-create/internal/pages"

	tea "charm.land/bubbletea/v2"
)

func (m TUIInterface) Init() tea.Cmd {
	return func() tea.Msg {
		return pages.ProjectNamePageMsg{}
	}
}
