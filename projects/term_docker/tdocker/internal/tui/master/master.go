package master

import (
	"tdocker/internal/components"
	"tdocker/internal/state"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type GotoPageMsg struct {
	Page int
}

const gap = "\n\n"

type (
	errMsg error
)

const (
	FOCUS_MENU = iota
	FOCUS_PAGE
)

type MasterModel struct {
	// viewport viewport.Model
	menu        components.MenuFrame
	senderStyle lipgloss.Style
	// messages    []string
	err    error
	npages map[int]tea.Model
	focus  int
	state  *state.State
}

// https://github.com/charmbracelet/bubbletea/tree/main/examples/chat
func NewMasterModel(menu components.MenuFrame) MasterModel {

	// ta := textarea.New()
	// ta.Placeholder = "Send a message..."
	// ta.Focus()

	// ta.Prompt = "┃ "
	// ta.CharLimit = 280

	// ta.SetWidth(30)
	// ta.SetHeight(3)

	// // Remove cursor line styling
	// ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	// ta.ShowLineNumbers = false

	// 	vp := viewport.New(30, 5)

	// 	vp.SetContent(`Welcome to the chat room!
	// Type a message and press Enter to send.`)

	// ta.KeyMap.InsertNewline.SetEnabled(false)

	app_state := &state.State{
		CurrentPage: state.PAGE_CONTAINERS,
	}

	return MasterModel{
		// viewport:    vp,
		menu: menu,
		// messages:    []string{},
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
		npages: map[int]tea.Model{
			state.PAGE_CONTAINERS: components.NewPageContainers(app_state),
			state.PAGE_IMAGES:     components.NewPageImages(),
			state.PAGE_VOLUMES:    components.NewPageVolumes(),
			state.PAGE_NETWORKS:   components.NewPageNetworks(),
		},
		focus: FOCUS_MENU,
		state: app_state,
	}
}

func (m MasterModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m MasterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = []tea.Cmd{}

	// vp, vpCmd := m.viewport.Update(msg)
	// cmds = append(cmds, vpCmd)
	// m.viewport = vp

	// menu
	m.menu.SetFocus(m.focus == FOCUS_MENU)
	if m.focus == FOCUS_MENU {
		menu, mCmd := m.menu.Update(msg)
		cmds = append(cmds, mCmd)
		m.menu = menu
	}

	// current page
	if m.focus == FOCUS_PAGE {
		// tea.Println("focus pages")
		p, pCmp := m.npages[m.state.CurrentPage].Update(msg)
		cmds = append(cmds, pCmp)
		m.npages[m.state.CurrentPage] = p
	}

	switch msg := msg.(type) {
	// case tea.WindowSizeMsg:
	// m.viewport.Width = msg.Width
	// m.textarea.SetWidth(msg.Width)
	// m.viewport.Height = msg.Height - m.textarea.Height() - lipgloss.Height(gap)
	// m.viewport.Height = msg.Height - m.textarea.Height() - (lipgloss.Height(gap) * 2) - 10

	// if len(m.messages) > 0 {
	// 	// Wrap content before setting it.
	// 	m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width).Render(strings.Join(m.messages, "\n")))
	// }
	// m.viewport.GotoBottom()

	// goto page event handling
	case GotoPageMsg:
		m.state.CurrentPage = msg.Page
		return m, nil

	// keys handling
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			// fmt.Println(m.textarea.Value())
			return m, tea.Quit
		// case tea.KeyCtrlN:

		// 	if m.pageIndex == 0 {
		// 		m.pageIndex = 1

		// 	} else {
		// 		m.pageIndex = 0
		// 	}
		// 	m.activePage = m.pages[m.pageIndex]
		// 	return m, nil
		case tea.KeyTab:
			if m.focus == FOCUS_MENU {
				m.focus = FOCUS_PAGE
			} else {
				m.focus = FOCUS_MENU
			}

			// ну вот как то не там где должен быть...
			// m.menu.SetFocus(m.focus == FOCUS_MENU)
			return m, nil

			// case tea.KeyEnter:
			// 	m.messages = append(m.messages, m.senderStyle.Render("You: ")+m.textarea.Value())
			// 	m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width).Render(strings.Join(m.messages, "\n")))
			// 	m.textarea.Reset()
			// 	m.viewport.GotoBottom()
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(cmds...)
}

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
)

func (m MasterModel) View() string {

	p_view := ""
	cp, ok := m.npages[m.state.CurrentPage]
	if ok {
		p_view = cp.View()
	}

	m_style := blurredStyle
	if m.focus == FOCUS_MENU {
		m_style = focusedStyle
	}

	p_style := blurredStyle
	if m.focus == FOCUS_PAGE {
		p_style = focusedStyle
	}

	// test columns
	// main_view := lipgloss.JoinHorizontal(lipgloss.Top, m.menu.View(), m.pages[m.pageIndex].View())
	main_view := lipgloss.JoinHorizontal(
		lipgloss.Top,
		m_style.Render(m.menu.View()),
		p_style.Render(p_view),
	)

	return lipgloss.JoinVertical(
		lipgloss.Top,
		main_view,
		gap,
		"Press Ctrl+C or Esc to exit",
	)

	// return fmt.Sprintf(
	// 	"%s%s%s",
	// 	// m.viewport.View(),
	// 	// "\nPress Ctrl+N to change page\n\n",
	// 	// m.menu.View(),
	// 	main_view,
	// 	gap,
	// 	gap,
	// 	// m.textarea.View(),
	// )
}
