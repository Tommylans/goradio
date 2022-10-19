package main

import (
	"log"
	"radio/player"
	"radio/ui"
)

func main() {
	radioPlayer := player.NewRadioPlayer()

	playerUi := ui.NewPlayerUi(radioPlayer)

	logger := log.New(playerUi.LogView(), "", 0)
	radioPlayer.SetLogger(logger)

	err := playerUi.Start()
	if err != nil {
		panic(err)
	}
}
