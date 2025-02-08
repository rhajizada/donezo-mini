package styles

import "github.com/charmbracelet/lipgloss"

var (
	App = lipgloss.NewStyle().
		Margin(1, 2)

	StatusMessage = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"})

	ErrorMessage = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#FB4A8A", Dark: "#FB4A8A"})

	Item = lipgloss.NewStyle().
		Padding(0, 0)

	Tags = lipgloss.NewStyle().
		AlignHorizontal(lipgloss.Center).
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4"))
)
