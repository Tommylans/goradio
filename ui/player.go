package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log"
	"os"
	"radio/channels"
	"radio/discord"
	"radio/radioplayer"
	"sync"
)

type PlayerUi struct {
	player      *radioplayer.RadioPlayer
	tracksTable *tview.Table
	logView     *tview.TextView

	logger *log.Logger

	currentChannel *channels.RadioChannel

	debugMode bool

	playLock sync.Mutex
}

func (p *PlayerUi) GetLogView() *tview.TextView {
	return p.logView
}

func (p *PlayerUi) SetLogger(logger *log.Logger) {
	p.logger = logger
}

func NewPlayerUi(player *radioplayer.RadioPlayer, debugMode bool) *PlayerUi {
	logView := tview.NewTextView()
	logView.SetTitle("Log")
	logView.SetTitleAlign(tview.AlignLeft)
	logView.SetBorder(true)

	return &PlayerUi{player: player, logView: logView, debugMode: debugMode}
}

func (p *PlayerUi) StartTui() error {
	app := tview.NewApplication()

	p.initTracksTable()
	p.setTracksTableData()

	mainView := tview.NewFlex()
	mainView.AddItem(p.tracksTable, 0, 1, true)
	if p.debugMode {
		mainView.AddItem(p.logView, 70, 1, false)
	}

	helpInfo := tview.NewTextView().
		SetTextColor(tcell.ColorBlue).
		SetText(" m: Mute, s: Stop, +/=: Increase volume, -: Decrease volume, 0: Reset volume, q: Quit")

	creditsText := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignRight).
		SetText("[green]Made with [red]❤  [green]by [gold]Tommylans")

	infoBox := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(helpInfo, 0, 1, false).
		AddItem(creditsText, 0, 1, false)

	view := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(mainView, 0, 1, true).
		AddItem(infoBox, 1, 1, false)

	appUi := app.SetRoot(view, true)

	mainView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
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

	p.currentChannel = channels.RadioChannels[row]

	go func() {
		err := discord.UpdateDiscordPresence(p.currentChannel)
		if err != nil {
			p.logger.Println("Error updating discord presence:", err)
		}
	}()
}

func (p *PlayerUi) setTracksTableData() {
	stations := channels.RadioChannels

	p.tracksTable.Clear()

	for i, station := range stations {
		p.tracksTable.SetCellSimple(i, 0, station.Name)
		p.tracksTable.SetCellSimple(i, 1, station.Url)
	}
}
