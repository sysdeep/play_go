package board

import (
	"fmt"
	"tcalendar/internal/calendar"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Selector struct {
	model *Model
}

func newSelector(model *Model) Selector {
	return Selector{
		model: model,
	}
}

func (s Selector) Init() tea.Cmd {
	// return tea.Batch(listenForActivity(a.sub), waitForActivity(a.sub))
	return nil
}
func (s Selector) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "j":
			s.model.PrevYear()
			return s, nil

		case "k":
			s.model.NextYear()
			return s, nil

		case "h":
			s.model.PrevMonth()
			return s, nil

		case "l":
			s.model.NextMonth()
			return s, nil

			// case "k":
			// 	a.currentPage = PAGE_KAFKA
			// 	return a, nil

		}

	}

	return s, nil
}

var sectionItemStyle = lipgloss.NewStyle().Width(10).AlignHorizontal(lipgloss.Center)
var arrowItemStyle = lipgloss.NewStyle().Width(3).AlignHorizontal(lipgloss.Center)
var spacerItemStyle = lipgloss.NewStyle().Width(2)

func (s Selector) View() string {

	monthName := calendar.MonthName(s.model.Month)

	return lipgloss.JoinHorizontal(lipgloss.Left,
		arrowItemStyle.Render("<j "),
		sectionItemStyle.Render(fmt.Sprintf("%d", s.model.Year)),
		arrowItemStyle.Render(" k>"),
		spacerItemStyle.Render(" "),
		arrowItemStyle.Render("<h "),
		sectionItemStyle.Render(monthName),
		arrowItemStyle.Render(" l>"),
	)
}
