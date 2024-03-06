package ui

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/tommylans/goradio/channels"
	"github.com/tommylans/goradio/discord"
)

func (p *PlayerUi) initTracksTable() {
	p.tracksTable = tview.NewTable()
	p.tracksTable.SetTitle("Radio Stations")
	p.tracksTable.SetTitleAlign(tview.AlignLeft)
	p.tracksTable.SetTitleColor(tcell.ColorYellow)
	p.tracksTable.SetBorder(true)
	p.tracksTable.SetSelectable(true, false)
	p.tracksTable.SetBorderPadding(0, 0, 1, 1)
	p.tracksTable.SetBorderColor(tcell.ColorBrown)

	p.tracksTable.SetSelectedFunc(func(row, column int) {
		p.playRow(row)
	})
}

func (p *PlayerUi) playRow(row int) {
	p.playLock.Lock()
	go func() {
		p.player.PlayRadioChannel(channels.RadioChannels[row])
		p.playLock.Unlock()
	}()

	p.renderTracksTableData()

	playingStation := p.tracksTable.GetCell(row, 0)
	playingStation.SetText("▶ " + playingStation.Text)
	playingStation.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorGreen))

	p.currentChannel = channels.RadioChannels[row]

	go func() {
		err := discord.UpdateDiscordPresence(p.currentChannel)
		if err != nil {
			log.Println("Error updating discord presence:", err)
		}
	}()
}

func (p *PlayerUi) renderTracksTableData() {
	p.tracksTable.Clear()

	for i, station := range channels.RadioChannels {
		p.tracksTable.SetCellSimple(i, 0, station.GetName())
		p.tracksTable.SetCellSimple(i, 1, station.GetLocation())
	}
}
