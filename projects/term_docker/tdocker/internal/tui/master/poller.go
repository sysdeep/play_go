package master

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// A message used to indicate that activity has occurred. In the real world (for
// example, chat) this would contain actual data.
type PollMsg struct{}

// TODO: see - https://github.com/charmbracelet/bubbletea/blob/main/examples/fullscreen/main.go
// Simulate a process that sends events at an irregular interval in real time.
// In this case, we'll send events on the channel at a random interval between
// 100 to 1000 milliseconds. As a command, Bubble Tea will run this
// asynchronously.
func listenForActivity(sub chan struct{}) tea.Cmd {
	return func() tea.Msg {
		for {
			time.Sleep(time.Second * 1)
			// time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)+100)) // nolint:gosec
			sub <- struct{}{}
		}
	}
}

// A command that waits for the activity on a channel.
func waitForActivity(sub chan struct{}) tea.Cmd {
	return func() tea.Msg {
		return PollMsg(<-sub)
	}
}
