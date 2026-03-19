package application

import (
	"tcalendar/internal/configuration"
	"tcalendar/internal/tui/about"
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

	router *Router

	brd tea.Model
}

const PAGE_CALENDAR = "calendar"
const PAGE_ABOUT = "about"

func New(conf configuration.Configuration) *Application {
	router := newRouter()
	router.Register(PAGE_CALENDAR, board.New())
	router.Register(PAGE_ABOUT, about.New(conf))

	return &Application{
		sub: make(chan struct{}),
		// brd:    board.New(),
		width:  minWidth,
		height: minHeight,
		router: router,
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

		case "a":
			a.router.GoTo(PAGE_ABOUT)
			return a, nil

		case "c":
			a.router.GoTo(PAGE_CALENDAR)
			return a, nil

		// These keys should exit the program.
		case "ctrl+c", "q":
			return a, tea.Quit
		}

	}

	var cmds = []tea.Cmd{}

	_, bCmd := a.router.MustPage().Update(msg)
	cmds = append(cmds, bCmd)

	return a, tea.Batch(cmds...)

}

func (a Application) View() string {

	pageView := a.router.MustPage().View()

	content := lipgloss.JoinVertical(lipgloss.Top, pageView) + "\n\n[a]bout [c]alendar [q]uit"

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
