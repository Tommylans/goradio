package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
	"radio/channels"
	"radio/player"
	"sync"
)

type PlayerUi struct {
	player      *player.RadioPlayer
	tracksTable *tview.Table
	logView     *tview.TextView

	debugMode bool

	playLock sync.Mutex
}

func (p *PlayerUi) LogView() *tview.TextView {
	return p.logView
}

func NewPlayerUi(player *player.RadioPlayer, debugMode bool) *PlayerUi {
	logView := tview.NewTextView()
	logView.SetTitle("Log")
	logView.SetTitleAlign(tview.AlignLeft)
	logView.SetBorder(true)

	return &PlayerUi{player: player, logView: logView, debugMode: debugMode}
}

func (p *PlayerUi) Start() error {
	flex := tview.NewFlex()
	flex.SetTitle("Radio Stations")
	flex.SetTitleAlign(tview.AlignLeft)
	flex.SetBorder(false)

	p.initTracksTable()
	p.setTracksTableData()

	flex.AddItem(p.tracksTable, 0, 1, true)
	if p.debugMode {
		flex.AddItem(p.logView, 70, 1, false)
	}

	appUi := tview.NewApplication().SetRoot(flex, true)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		key := string(event.Rune())

		fmt.Fprintln(p.logView, "Keypress:", key, event.Key(), event.Rune(), event.Name())

		switch key {
		case "s":
			p.player.Stop()
			p.setTracksTableData()
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
		case "q":
			appUi.Stop()
			p.player.Close()
			os.Exit(0)
		}

		return event
	})

	err := appUi.Run()
	if err != nil {
		return err
	}

	return nil
}

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
		p.player.PlayChannel(channels.RadioChannels[row])
		p.playLock.Unlock()
	}()

	p.setTracksTableData()

	playingStation := p.tracksTable.GetCell(row, 0)
	playingStation.SetText("▶ " + playingStation.Text)
	playingStation.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorGreen))
}

func (p *PlayerUi) setTracksTableData() {
	stations := channels.RadioChannels

	p.tracksTable.Clear()

	for i, station := range stations {
		p.tracksTable.SetCellSimple(i, 0, station.Name)
		p.tracksTable.SetCellSimple(i, 1, station.Url)
	}
}
