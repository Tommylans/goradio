package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"radio/channels"
	"radio/radioplayer"
	"sync"
)

type PlayerUi struct {
	player      *radioplayer.RadioPlayer
	tracksTable *tview.Table
	logView     *tview.TextView

	currentChannel *channels.RadioChannel

	debugMode bool

	playLock sync.Mutex
	appUi    *tview.Application
}

func NewPlayerUi(player *radioplayer.RadioPlayer, debugMode bool) *PlayerUi {
	logView := tview.NewTextView()
	logView.SetTitle("Log")
	logView.SetTitleAlign(tview.AlignLeft)
	logView.SetBorder(true)

	return &PlayerUi{
		player:    player,
		logView:   logView,
		debugMode: debugMode,
	}
}

func (p *PlayerUi) GetLogView() *tview.TextView {
	return p.logView
}

func (p *PlayerUi) StartTui() error {
	app := tview.NewApplication()

	p.initTracksTable()
	p.renderTracksTableData()

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

	p.appUi = app.SetRoot(view, true)

	mainView.SetInputCapture(p.playerKeyBindHandler())
	p.appUi.SetInputCapture(p.applicationKeyBindHandler())

	err := p.appUi.Run()
	if err != nil {
		return err
	}

	return nil
}
