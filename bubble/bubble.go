package bubble

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tommylans/goradio/channels"
)

type Bubble struct {
	program *tea.Program
}

func NewBubble() *Bubble {
	program := tea.NewProgram(channelsList(channels.RadioChannels), tea.WithAltScreen())
	return &Bubble{
		program: program,
	}
}

func (b *Bubble) Run() error {
	_, err := b.program.Run()
	return err
}
