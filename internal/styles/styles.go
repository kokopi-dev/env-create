package styles

import "charm.land/lipgloss/v2"

var (
	LabelStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	HintStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true)
	ConfirmStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
)
