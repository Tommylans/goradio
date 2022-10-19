package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"radio/channels"
	"radio/player"
)

type PlayerUi struct {
	player      *player.RadioPlayer
	tracksTable *tview.Table
	logView     *tview.TextView
}

func (p *PlayerUi) LogView() *tview.TextView {
	return p.logView
}

func NewPlayerUi(player *player.RadioPlayer) *PlayerUi {
	logView := tview.NewTextView()
	logView.SetTitle("Log")
	logView.SetTitleAlign(tview.AlignLeft)
	logView.SetBorder(true)

	return &PlayerUi{player: player, logView: logView}
}

func (p *PlayerUi) Start() error {
	flex := tview.NewFlex()
	flex.SetTitle("Radio Stations")
	flex.SetTitleAlign(tview.AlignLeft)
	flex.SetBorder(false)

	p.initTracksTable()
	p.setTracksTableData()

	flex.AddItem(p.tracksTable, 70, 0, true)
	flex.AddItem(p.logView, 50, 1, false)

	err := tview.NewApplication().SetRoot(flex, false).Run()
	if err != nil {
		return err
	}

	return nil
}

func (p *PlayerUi) initTracksTable() {
	p.tracksTable = tview.NewTable()
	p.tracksTable.SetTitle("Radio Stations")
	p.tracksTable.SetBorder(true)
	p.tracksTable.SetSelectable(true, false)
	p.tracksTable.SetSelectedStyle(tcell.StyleDefault.Underline(true))
	p.tracksTable.SetBorderPadding(0, 0, 1, 1)

	p.tracksTable.SetSelectedFunc(func(row, column int) {
		go p.player.PlayChannel(channels.RadioChannels[row])

		p.setTracksTableData()

		playingStation := p.tracksTable.GetCell(row, 0)
		playingStation.SetText("▶ " + playingStation.Text)
	})
}

func (p *PlayerUi) setTracksTableData() {
	stations := channels.RadioChannels

	p.tracksTable.Clear()

	for i, station := range stations {
		p.tracksTable.SetCellSimple(i, 0, station.Name)
		p.tracksTable.SetCellSimple(i, 1, station.Url)
	}
}
