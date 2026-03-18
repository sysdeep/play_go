package board

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Board struct {
	model    *Model
	selector tea.Model
	days     tea.Model
	todayBar tea.Model
}

func New() *Board {
	model := newModel()
	return &Board{
		model:    model,
		selector: newSelector(model),
		days:     newDays(model),
		todayBar: newTodayBar(),
	}
}

func (b Board) Init() tea.Cmd {
	return nil
}

func (b Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	newSelector, _ := b.selector.Update(msg)
	b.selector = newSelector

	newDays, _ := b.days.Update(msg)
	b.days = newDays

	return b, nil
}

func (b Board) View() string {

	return lipgloss.JoinVertical(lipgloss.Bottom,
		b.todayBar.View(),
		" ",
		b.selector.View(),
		" ",
		b.days.View(),
	)
}
