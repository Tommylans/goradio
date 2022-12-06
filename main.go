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
	playerUi.SetLogger(logger)
	radioPlayer.SetLogger(logger)

	go playerUi.InitDiscordRichPresence("1049721387922751528")

	err := playerUi.Start()
	if err != nil {
		panic(err)
	}
}
