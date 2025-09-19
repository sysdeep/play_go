package components

import (
	"context"
	"strings"
	"tdocker/internal/tui"
	"tdocker/internal/utils"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type PageContainers struct {
	t            table.Model
	focused      bool
	dockerCli    *client.Client
	pageGeometry *PageGeometry
}

func NewPageContainers(dockerCli *client.Client, pageGeometry *PageGeometry) PageContainers {

	columns := []table.Column{
		{Title: "Status", Width: 20},
		{Title: "ID", Width: 20},
		{Title: "Name", Width: 20},
		{Title: "Image", Width: 30},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true), table.WithHeight(7))

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
		t:            t,
		focused:      false,
		dockerCli:    dockerCli,
		pageGeometry: pageGeometry,
	}
}

func (p PageContainers) Init() tea.Cmd {
	p.updateList()
	return nil
}

func (p PageContainers) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// проверка на фокус
	switch msg := msg.(type) {

	// case tui.GotoPageMsg:
	// 	return p, tea.Println(msg.Page)
	// 	if msg.Page == tui.PAGE_CONTAINERS {
	// 		p.updateList()
	// 		// return p, nil
	// 	}

	// если событие обновления странички - обновляем
	case tui.NeedRefreshMsg:
		// return p, tea.Println("need")
		if msg.Page == tui.PAGE_CONTAINERS {
			p.updateList()
			return p, nil
		}

	case tui.FocusMsg:
		oldFocus := p.focused
		p.focused = msg.Focus == tui.FOCUS_PAGE

		// при получении фокуса - обновляем таблицу
		if p.focused != oldFocus && p.focused {
			p.updateList()
		}
		return p, nil
	}

	// если не в фокусе - события не обрабатываем
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

	borderStyle := MakeFocusedBorder(p.focused)

	screen := p.t.View()
	return borderStyle.Render(
		lipgloss.NewStyle().Height(p.pageGeometry.MaxHeight).Width(p.pageGeometry.MaxWidth).Render(screen))
}

func (p *PageContainers) updateList() {
	raw_containers, _ := p.dockerCli.ContainerList(context.Background(), container.ListOptions{All: true})

	rows := []table.Row{}
	for _, c := range raw_containers {

		shortID := utils.ShortID(c.ID)

		rows = append(rows, []string{
			c.Status,                    // Status
			shortID,                     // ID
			strings.Join(c.Names, ", "), // Name
			c.Image,                     // Image
		})
	}
	p.t.SetRows(rows)

}
