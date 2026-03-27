package tui

import (
	"env-create/internal/pages"

	tea "charm.land/bubbletea/v2"
)

// initial entrypoint state/page
func (m TUIInterface) Init() tea.Cmd {
	return func() tea.Msg {
		return pages.HomePageMsg{}
	}
}
