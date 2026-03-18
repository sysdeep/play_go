package board

import "github.com/charmbracelet/lipgloss"

const dayWidth = 5
const daysInWeek = 7

var rowStyle = lipgloss.NewStyle().Width(daysInWeek * dayWidth).AlignHorizontal(lipgloss.Center)
var dayStyle = lipgloss.NewStyle().Width(dayWidth).AlignHorizontal(lipgloss.Center)
