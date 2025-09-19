package components

import (
	"tdocker/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type PageNetworks struct {
	focused      bool
	pageGeometry *PageGeometry
}

func NewPageNetworks(pageGeometry *PageGeometry) PageNetworks {
	return PageNetworks{
		focused:      false,
		pageGeometry: pageGeometry,
	}
}

func (p PageNetworks) Init() tea.Cmd {
	return nil
}

func (p PageNetworks) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tui.FocusMsg:
		p.focused = msg.Focus == tui.FOCUS_PAGE
		return p, nil
	}

	return p, nil
}

func (p PageNetworks) View() string {
	borderStyle := MakeFocusedBorder(p.focused)
	return borderStyle.Render(
		lipgloss.NewStyle().Width(p.pageGeometry.MaxWidth).Height(p.pageGeometry.MaxHeight).Render("This is network page"),
	)
}
