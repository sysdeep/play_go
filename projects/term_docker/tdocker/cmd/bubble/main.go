package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {

	// m := makeMainMenu()

	p := tea.NewProgram(
		// newSimplePage("This app is under construction"),
		// m,
		makeMainModel(),
	)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

func makeMainMenu() menu {
	return menu{
		options: []menuItem{
			{
				text:    "Containers",
				onPress: func() tea.Msg { return toggleCasingMsg{} },
			},
			{
				text:    "Images",
				onPress: func() tea.Msg { return toggleCasingMsg{} },
			},
			{
				text:    "Volumes",
				onPress: func() tea.Msg { return toggleCasingMsg{} },
			},
			{
				text:    "Networks",
				onPress: func() tea.Msg { return toggleCasingMsg{} },
			},
			{
				text:    "exit",
				onPress: func() tea.Msg { return exitMsg{} },
			},
		},
	}
}

// new
const gap = "\n\n"

type (
	errMsg error
)

type mainModel struct {
	viewport    viewport.Model
	textarea    textarea.Model
	menu        menu
	senderStyle lipgloss.Style
	messages    []string
	err         error
	pages       []tea.Model
	pageIndex   int
}

// https://github.com/charmbracelet/bubbletea/tree/main/examples/chat
func makeMainModel() mainModel {

	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 280

	ta.SetWidth(30)
	ta.SetHeight(3)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(30, 5)

	vp.SetContent(`Welcome to the chat room!
Type a message and press Enter to send.`)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return mainModel{
		textarea:    ta,
		viewport:    vp,
		menu:        makeMainMenu(),
		messages:    []string{},
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
		pages: []tea.Model{
			pageContainers{},
			pageImages{},
		},
		pageIndex: 0,
	}
}

func (m mainModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
		mCmd  tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)
	m.menu, mCmd = m.menu.Update(msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// m.viewport.Width = msg.Width
		// m.textarea.SetWidth(msg.Width)
		// m.viewport.Height = msg.Height - m.textarea.Height() - lipgloss.Height(gap)
		// m.viewport.Height = msg.Height - m.textarea.Height() - (lipgloss.Height(gap) * 2) - 10

		if len(m.messages) > 0 {
			// Wrap content before setting it.
			m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width).Render(strings.Join(m.messages, "\n")))
		}
		m.viewport.GotoBottom()
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			fmt.Println(m.textarea.Value())
			return m, tea.Quit
		case tea.KeyCtrlN:

			if m.pageIndex == 0 {
				m.pageIndex = 1
			} else {
				m.pageIndex = 0
			}
			return m, nil
		case tea.KeyEnter:
			m.messages = append(m.messages, m.senderStyle.Render("You: ")+m.textarea.Value())
			m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width).Render(strings.Join(m.messages, "\n")))
			m.textarea.Reset()
			m.viewport.GotoBottom()
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd, mCmd)
}

func (m mainModel) View() string {

	// test columns
	a := lipgloss.JoinHorizontal(lipgloss.Top, m.menu.View(), m.pages[m.pageIndex].View())

	return fmt.Sprintf(
		"%s%s%s%s%s",
		m.viewport.View(),
		"\nPress Ctrl+N to change page\n\n",
		// m.menu.View(),
		a,
		gap,
		m.textarea.View(),
	)
}
