package components

import (
	"context"
	"fmt"
	"strings"
	"tdocker/internal/tui"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// var baseStyle = lipgloss.NewStyle().
// 	BorderStyle(lipgloss.NormalBorder()).
// 	BorderForeground(lipgloss.Color("240"))

type PageContainers struct {
	// state *state.State
	t          table.Model
	focused    bool
	dockerCli  *client.Client
	conteiners []container.Summary
	MaxWidth   int
	MaxHeight  int
}

func NewPageContainers(dockerCli *client.Client) PageContainers {

	// raw_containers, _ := dockerCli.ContainerList(context.Background(), container.ListOptions{All: true})

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
		// state: st,
		t:         t,
		focused:   false,
		dockerCli: dockerCli,
		MaxWidth:  70,
		MaxHeight: 30,
		// conteiners: raw_containers,
	}
}

func (p PageContainers) Init() tea.Cmd {
	return nil
}

func (p PageContainers) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// tea.Println("update containes")
	// if p.state.CurrentPage == state.PAGE_CONTAINERS {
	// 	p.t.Focus()
	// }

	switch msg := msg.(type) {

	case tui.FocusMsg:
		p.focused = msg.Focus == tui.FOCUS_PAGE
		return p, nil
	}

	if !p.focused {
		return p, nil
	}

	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:

		// switch msg.Type {

		// case tea.KeyCtrlN:
		// 	p.t.Focus()
		// }

		switch msg.String() {

		case "r":
			p.updateList()
			return p, nil
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

func (p PageContainers) View() string {
	// return "This is containers page"

	wStyle := lipgloss.NewStyle().Width(p.MaxWidth)
	// fmt.Println(p.MaxWidth)

	rows := p.t.Rows()
	for _, c := range p.conteiners {
		rows = append(rows, []string{c.Status, c.ID, c.Names[0], "-"})
	}
	p.t.SetRows(rows)

	borderStyle := MakeFocusedBorder(p.focused)

	cs := []string{}
	for _, c := range p.conteiners {
		cs = append(cs, wStyle.Render(c.ID))
	}

	ccs := strings.Join(cs, "\n")

	v := lipgloss.JoinVertical(lipgloss.Bottom, p.t.View(), ccs,
		wStyle.Render(fmt.Sprintf("len: %d", len(rows))),
	)

	// return borderStyle.Render(p.t.View())
	return borderStyle.Render(lipgloss.NewStyle().Height(p.MaxHeight).Render(v))
}

func (p *PageContainers) updateList() {
	raw_containers, _ := p.dockerCli.ContainerList(context.Background(), container.ListOptions{All: true})
	p.conteiners = raw_containers
}
