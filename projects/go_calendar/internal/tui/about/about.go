package about

import (
	"tcalendar/internal/configuration"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type About struct {
	conf configuration.Configuration
}

func New(conf configuration.Configuration) About {
	return About{
		conf: conf,
	}
}

func (a About) Init() tea.Cmd {
	return nil
}

func (a About) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "esc":
			// pass

		}
	}
	return a, nil
}

func (a About) View() string {

	// pageView := a.router.MustPage().View()

	// content := lipgloss.JoinVertical(lipgloss.Top, pageView)

	// return lipgloss.PlaceVertical(
	// 	a.height,
	// 	lipgloss.Center,
	// 	lipgloss.PlaceHorizontal(
	// 		a.width,
	// 		lipgloss.Center,
	// 		content,
	// 	),
	// )

	return lipgloss.JoinVertical(
		lipgloss.Center,
		"TCalendar - simple terminal calendar\n",
		"version: "+a.conf.AppVersion,
	)
}
