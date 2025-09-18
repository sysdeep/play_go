package components

import "github.com/charmbracelet/lipgloss"

func MakeFocusedBorder(isFocus bool) lipgloss.Style {
	color := "240"
	if isFocus {
		color = "205"
	}

	return lipgloss.NewStyle().
		// BorderStyle(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(color))

}
