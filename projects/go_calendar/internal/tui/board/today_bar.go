package board

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TodayBar struct{}

func newTodayBar() TodayBar {
	return TodayBar{}
}

func (t TodayBar) Init() tea.Cmd {
	return nil
}

func (t TodayBar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	return t, nil
}

func (t TodayBar) View() string {

	now := time.Now()
	result := now.Format(time.DateTime)
	return rowStyle.Render(result)
}
