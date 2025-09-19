package tui

import tea "github.com/charmbracelet/bubbletea"

const (
	PAGE_CONTAINERS = 1
	PAGE_IMAGES     = 2
	PAGE_VOLUMES    = 3
	PAGE_NETWORKS   = 4
)

type GotoPageMsg struct {
	Page int
}

func MakeGotoPageMsg(page int) tea.Cmd {
	return func() tea.Msg {
		return GotoPageMsg{
			Page: page,
		}
	}
}

type NeedRefreshMsg struct {
	Page int
}
