package ui

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func (p *PlayerUi) playerKeyBindHandler() func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		key := string(event.Rune())

		log.Println("Key pressed:", key)

		switch key {
		case "s":
			p.player.Stop()
			p.renderTracksTableData()
			break
		case "-":
			p.player.DecreaseVolume()
			break
		case "+":
		case "=":
			p.player.IncreaseVolume()
			break
		case "0":
			p.player.ResetVolume()
			break
		case "m":
			p.player.Mute()
			break
		case " ":
			row, _ := p.tracksTable.GetSelection()
			p.playRow(row)
			break
		}

		return event
	}
}

func (p *PlayerUi) applicationKeyBindHandler() func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		key := string(event.Rune())

		if key == "q" {
			p.appUi.Stop()
			p.player.Close()
			os.Exit(0)
		}

		return event
	}
}
