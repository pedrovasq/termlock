package styles

import lip "github.com/charmbracelet/lipgloss"

var (
	HighLight	= lip.NewStyle().Bold(true).Foreground(lip.Color("2"))
	Normal		= lip.NewStyle().Foreground(lip.Color("15"))
	Bright		= lip.NewStyle().Bold(true).Foreground(lip.Color("#FFFFFF"))
	Faded		= lip.NewStyle().Foreground(lip.Color("7"))
	Password	= lip.NewStyle().Foreground(lip.Color("1")).Bold(true)
	Link		= lip.NewStyle().Foreground(lip.Color("6")).Bold(true)
	Error		= lip.NewStyle().Foreground(lip.Color("1"))
	Title		= lip.NewStyle().Bold(true).Underline(true).Foreground(lip.Color("5"))

	Border = lip.NewStyle().
		Border(lip.RoundedBorder()).
		Padding(1, 2).
		BorderForeground(lip.Color("15"))
)
