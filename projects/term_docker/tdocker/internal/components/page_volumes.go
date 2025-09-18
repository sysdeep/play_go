package components

import (
	"tdocker/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

type PageVolumes struct {
	focused bool
}

func NewPageVolumes() PageVolumes {
	return PageVolumes{
		focused: false,
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
	return borderStyle.Render("This is volumes page")
}
