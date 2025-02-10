package components

import (
	"tdocker/internal/state"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type PageContainers struct {
	state *state.State
	t     table.Model
}

func NewPageContainers(st *state.State) PageContainers {

	columns := []table.Column{
		{Title: "Rank", Width: 4},
		{Title: "City", Width: 10},
		{Title: "Country", Width: 10},
		{Title: "Population", Width: 10},
	}

	rows := []table.Row{
		{"1", "Tokyo", "Japan", "37,274,000"},
		{"2", "Delhi", "India", "32,065,760"},
		{"3", "Shanghai", "China", "28,516,904"},
		{"4", "Dhaka", "Bangladesh", "22,478,116"},
		{"5", "São Paulo", "Brazil", "22,429,800"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows), table.WithFocused(true), table.WithHeight(7))

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return PageContainers{
		state: st,
		t:     t,
	}
}

func (p PageContainers) Init() tea.Cmd {
	return nil
}

func (p PageContainers) View() string {
	// return "This is containers page"

	return baseStyle.Render(p.t.View()) + "\n"
}

func (p PageContainers) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// tea.Println("update containes")
	// if p.state.CurrentPage == state.PAGE_CONTAINERS {
	// 	p.t.Focus()
	// }

	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:

		// switch msg.Type {

		// case tea.KeyCtrlN:
		// 	p.t.Focus()
		// }

		switch msg.String() {

		case "esc":
			if p.t.Focused() {
				p.t.Blur()
			} else {
				p.t.Focus()
			}
		// case "q", "ctrl+c":
		// 	return p, tea.Quit
		case "enter":
			return p, tea.Batch(
				tea.Printf("Let's go to %s!", p.t.SelectedRow()[1]),
			)
		}
	}
	p.t, cmd = p.t.Update(msg)
	return p, cmd
}

func (p PageContainers) Focus() {
	p.t.Focus()
}
