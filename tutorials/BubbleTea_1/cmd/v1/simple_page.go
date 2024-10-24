package main

import (
	"fmt"
	"strings"
)

import (
	tea "github.com/charmbracelet/bubbletea"
)

// model
type SimplePage struct {
	text string
}

func NewSimplePage(text string) *SimplePage {
	return &SimplePage{
		text: text,
	}
}

func (s SimplePage) Init() tea.Cmd {
	return nil
}

// view

func (s SimplePage) View() string {
	text_len := len(s.text)
	bar := strings.Repeat("*", text_len+4)
	return fmt.Sprintf(
		"%s\n %s \n%s\n\nPress Ctrl+C to exit",
		bar, s.text, bar,
	)
}

// update
func (s SimplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return s, tea.Quit
		}
	}
	return s, nil
}
