package bubble

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tommylans/goradio/channels"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type radioChannelsList struct {
	list list.Model
}

func (r *radioChannelsList) Init() tea.Cmd {
	return nil
}

func (r *radioChannelsList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return r, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		r.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	r.list, cmd = r.list.Update(msg)
	return r, cmd
}

func (r *radioChannelsList) View() string {
	return docStyle.Render(r.list.View())
}

func channelsList(stations []channels.RadioStation) *radioChannelsList {
	var items []list.Item
	for _, station := range stations {
		items = append(items, &RadioChannelItem{station})
	}

	return &radioChannelsList{
		list: list.New(items, list.NewDefaultDelegate(), 0, 0),
	}
}

type RadioChannelItem struct {
	channels.RadioStation
}

func (r *RadioChannelItem) Title() string {
	return r.GetName()
}

func (r *RadioChannelItem) Description() string {
	return r.GetLocation()
}

func (r *RadioChannelItem) FilterValue() string {
	return r.GetName()
}
