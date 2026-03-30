package styles

import "charm.land/lipgloss/v2"

var (
	LabelStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	HintStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true)
	ConfirmStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))

	// Page layout
	FullScreenStyle = lipgloss.NewStyle().Padding(0)

	// Card / box — border only, no padding
	CardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			Width(52)

	// Inner content area with padding
	CardInnerStyle = lipgloss.NewStyle().
			Padding(1, 3)

	CardTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("86")).
			Bold(true).
			MarginBottom(1)

	CardSubtitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("245")).
				MarginBottom(1)

	InputLabelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245")).
			MarginBottom(1)

	// Footer — full card inner width, no side padding
	FooterStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			BorderTop(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("237")).
			Padding(0, 1).
			Width(50)

	FooterKeyStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("86")).Bold(true)
	FooterSepStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("237"))
	FooterDescStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("243"))

	InputLabelActiveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("86")).MarginBottom(1)
	ErrorStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
)
