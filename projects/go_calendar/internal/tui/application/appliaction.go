package application

import (
	"tcalendar/internal/tui/board"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const minHeight = 12
const minWidth = 5 * 7

type Application struct {
	width  int
	height int

	sub chan struct{}

	brd tea.Model
}

func New() *Application {
	return &Application{
		sub:    make(chan struct{}),
		brd:    board.New(),
		width:  minWidth,
		height: minHeight,
	}
}

func (a Application) Init() tea.Cmd {
	return tea.Batch(listenForActivity(a.sub), waitForActivity(a.sub))
}

func (a Application) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	// loop activity
	case PollMsg:
		return a, waitForActivity(a.sub) // wait for next event

	case tea.WindowSizeMsg:
		if msg.Height < minHeight {
			a.height = minHeight
		} else {
			a.width = msg.Width
		}
		if msg.Width < minWidth {
			a.width = minWidth
		} else {
			a.height = msg.Height
		}
		return a, nil

	// case core.GotoPageMsg:
	// 	if msg.Page == core.PAGE_BACK {
	// 		a.pagesStack = a.pagesStack[:len(a.pagesStack)-1]
	// 	} else {
	// 		a.pagesStack = append(a.pagesStack, msg.Page)
	// 	}
	// 	a.page = a.pagesStack[len(a.pagesStack)-1]
	// 	return a, nil

	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {

		// case "d":
		// 	a.currentPage = PAGE_DASHBOARD
		// 	return a, nil

		// case "a":
		// 	a.currentPage = PAGE_AM
		// 	return a, nil

		// case "k":
		// 	a.currentPage = PAGE_KAFKA
		// 	return a, nil

		// These keys should exit the program.
		case "ctrl+c", "q":
			return a, tea.Quit
		}

	}

	var cmds = []tea.Cmd{}

	newBoard, bCmd := a.brd.Update(msg)
	a.brd = newBoard
	cmds = append(cmds, bCmd)

	// // current page
	// p, pCmp := a.pages[a.page].Update(msg)
	// cmds = append(cmds, pCmp)
	// a.pages[a.page] = p

	// all pages
	// for k, p := range a.pages {
	// 	pp, ppCmd := p.Update(msg)
	// 	cmds = append(cmds, ppCmd)
	// 	a.pages[k] = pp
	// }

	return a, tea.Batch(cmds...)

}

func (a Application) View() string {

	// full_result := lipgloss.PlaceVertical(
	// 	a.maxHeight-2,
	// 	lipgloss.Center,
	// 	lipgloss.PlaceHorizontal(
	// 		a.maxWidth,
	// 		lipgloss.Center,
	// 		a.pages[a.page].View(),
	// 		lipgloss.WithWhitespaceBackground(lipgloss.Color("#3458eb")),
	// 	),
	// )

	// var pp = []string{}
	// for _, p := range a.pagesStack {
	// 	pp = append(pp, strconv.Itoa(p))
	// }

	// return lipgloss.NewStyle().
	// 	Width(a.maxWidth).
	// 	Height(a.maxHeight).
	// 	Background(lipgloss.Color("#3458eb")).
	// 	Render(lipgloss.JoinVertical(
	// 		lipgloss.Left,
	// 		full_result,
	// 		strings.Join(pp, ":"),
	// 	))
	content := lipgloss.JoinVertical(lipgloss.Top, a.brd.View())

	// return lipgloss.NewStyle().Width(300).Height(200).AlignHorizontal(lipgloss.Center).AlignVertical(lipgloss.Center).Render(content)

	return lipgloss.PlaceVertical(
		a.height,
		lipgloss.Center,
		lipgloss.PlaceHorizontal(
			a.width,
			lipgloss.Center,
			content,
		),
	)
}
