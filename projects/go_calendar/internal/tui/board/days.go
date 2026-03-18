package board

import (
	"fmt"
	"tcalendar/internal/calendar"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Days struct {
	model *Model
}

func newDays(model *Model) Days {
	return Days{
		model: model,
	}
}

func (d Days) Init() tea.Cmd {
	return nil
}

func (d Days) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// switch msg := msg.(type) {
	// case tea.KeyMsg:
	// 	switch msg.String() {

	// 	}

	// }

	return d, nil
}

func (d Days) View() string {

	daysView := [][]string{}

	dataset := []DayCell{}

	daysIndex := 0
	rowIndex := 0
	for {

		if len(daysView) < rowIndex+1 {
			daysView = append(daysView, []string{})
		}

		for _, w := range calendar.Weekdays {

			day := d.model.Days[daysIndex]
			if day.Weekday == w {
				dataset = append(dataset, day)
				dayValue := fmt.Sprintf("%d", day.Day)
				var dayStrValue string
				if d.model.IsCurrent(day) {
					dayStrValue = dayStyle.Bold(true).Render(">" + dayValue + "<")
				} else {
					dayStrValue = dayStyle.Render(dayValue)
				}
				daysView[rowIndex] = append(daysView[rowIndex], dayStrValue)
				if daysIndex+1 < len(d.model.Days) {

					daysIndex += 1
				}
			} else {
				dataset = append(dataset, DayCell{Day: 0})
				daysView[rowIndex] = append(daysView[rowIndex], dayStyle.Render(" "))
			}
		}
		rowIndex += 1
		if daysIndex+1 >= len(d.model.Days) {
			break
		}
	}

	rows := []string{}
	for _, row := range daysView {
		rowString := lipgloss.JoinHorizontal(lipgloss.Left, row...)
		rows = append(rows, rowString)
	}

	if len(rows) < 6 {
		for range 6 - len(rows) {
			rows = append(rows, " ")
		}
	}

	return lipgloss.JoinVertical(lipgloss.Top,
		lipgloss.JoinHorizontal(lipgloss.Left,
			dayStyle.Render("пн"),
			dayStyle.Render("вт"),
			dayStyle.Render("ср"),
			dayStyle.Render("чт"),
			dayStyle.Render("пт"),
			dayStyle.Render("сб"),
			dayStyle.Render("вс"),
		),
		lipgloss.JoinVertical(lipgloss.Top,
			rows...,
		),
	)
}
