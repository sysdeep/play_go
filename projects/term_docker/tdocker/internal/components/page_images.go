package components

import (
	"context"
	"tdocker/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

type PageImages struct {
	selectedIndex int
	options       []string
	focused       bool
	dockerCli     *client.Client
	images        []image.Summary
	pageGeometry  *PageGeometry
}

func NewPageImages(dockerCli *client.Client, pageGeometry *PageGeometry) PageImages {
	p := PageImages{
		selectedIndex: 0,
		options:       []string{"astra", "ubuntu"},
		focused:       false,
		dockerCli:     dockerCli,
		pageGeometry:  pageGeometry,
	}

	p.updateList()

	return p
}

func (p PageImages) Init() tea.Cmd {
	return nil
}

func (p PageImages) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	// case toggleCasingMsg:
	// 	return m.toggleSelectedItemCase(), nil
	// case exitMsg:
	// 	return m, tea.Quit

	case tui.FocusMsg:
		// fmt.Print(msg)
		p.focused = msg.Focus == tui.FOCUS_PAGE
		return p, nil
	}

	if !p.focused {
		return p, nil
	}

	switch msg := msg.(type) {
	// case toggleCasingMsg:
	// 	return m.toggleSelectedItemCase(), nil
	// case exitMsg:
	// 	return m, tea.Quit

	case tea.KeyMsg:
		switch msg.String() {

		case "r":
			p.updateList()
			return p, nil

		// case "ctrl+c":
		// 	return m, tea.Quit
		case "down", "right", "up", "left":
			if p.focused {
				p.moveCursor(msg.String())
			}
			return p, nil
		}
	}
	return p, nil
}

func (p PageImages) View() string {
	// var options []string
	// for i, o := range p.options {
	// 	if i == p.selectedIndex {
	// 		options = append(options, fmt.Sprintf("-> %s", o))
	// 	} else {
	// 		options = append(options, fmt.Sprintf("   %s", o))
	// 	}
	// }

	// body := strings.Join(options, "\n")

	imagesStr := []string{
		// body,
	}
	for idx, img := range p.images {
		if len(img.RepoTags) == 0 {
			continue
		}
		rowStr := img.RepoTags[0]
		preStr := "  "
		if idx == p.selectedIndex {
			preStr = "> "
		}
		imagesStr = append(imagesStr, preStr+rowStr)
	}

	body := lipgloss.JoinVertical(lipgloss.Left, imagesStr...)

	borderStyle := MakeFocusedBorder(p.focused)

	return borderStyle.Render(
		lipgloss.NewStyle().Width(p.pageGeometry.MaxWidth).Height(p.pageGeometry.MaxHeight).Render(body),
	)

}

func (p *PageImages) moveCursor(msg string) {
	switch msg {
	case "up", "left":
		p.selectedIndex--
	case "down", "right":
		p.selectedIndex++
	default:
		// do nothing
	}

	optCount := len(p.images)
	p.selectedIndex = (p.selectedIndex + optCount) % optCount
}

func (p *PageImages) updateList() {
	p.images, _ = p.dockerCli.ImageList(context.Background(), image.ListOptions{All: true})
}
