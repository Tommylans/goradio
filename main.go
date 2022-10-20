package main

import (
	"flag"
	"log"
	"radio/player"
	"radio/ui"
)

var (
	debug = flag.Bool("debug", false, "Will show the logging panel")
)

func main() {
	flag.Parse()

	radioPlayer := player.NewRadioPlayer()

	playerUi := ui.NewPlayerUi(radioPlayer, *debug)

	logger := log.New(playerUi.LogView(), "", 0)
	radioPlayer.SetLogger(logger)

	err := playerUi.Start()
	if err != nil {
		panic(err)
	}
}
