package components

import (
	"tdocker/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

type PageNetworks struct {
	focused bool
}

func NewPageNetworks() PageNetworks {
	return PageNetworks{
		focused: false,
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
	return borderStyle.Render("This is networks page")
}
