package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	FOCUS_MENU = "menu"
	FOCUS_PAGE = "page"
)

type FocusMsg struct {
	Focus string
}

func MakeFocusMsg(value string) tea.Cmd {
	return func() tea.Msg {
		return FocusMsg{Focus: value}
	}
}
