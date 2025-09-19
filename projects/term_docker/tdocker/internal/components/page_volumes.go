package components

import (
	"tdocker/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type PageVolumes struct {
	focused      bool
	pageGeometry *PageGeometry
}

func NewPageVolumes(pageGeometry *PageGeometry) PageVolumes {
	return PageVolumes{
		focused:      false,
		pageGeometry: pageGeometry,
	}
}

func (p PageVolumes) Init() tea.Cmd {
	return nil
}

func (p PageVolumes) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tui.FocusMsg:
		p.focused = msg.Focus == tui.FOCUS_PAGE
		return p, nil
	}

	return p, nil
}

func (p PageVolumes) View() string {
	borderStyle := MakeFocusedBorder(p.focused)
	return borderStyle.Render(
		lipgloss.NewStyle().Width(p.pageGeometry.MaxWidth).Height(p.pageGeometry.MaxHeight).Render("This is volumes page"),
	)
}
