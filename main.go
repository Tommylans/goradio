package main

import (
	"flag"
	"log"
	"radio/discord"
	"radio/radioplayer"
	"radio/ui"
)

var (
	debug = flag.Bool("debug", false, "Will show the logging panel")
)

const (
	// DiscordClientId is the token used to connect to Discord
	DiscordClientId = "1049721387922751528"
)

func main() {
	flag.Parse()

	radioPlayer := radioplayer.NewRadioPlayer()

	playerUi := ui.NewPlayerUi(radioPlayer, *debug)
	log.SetOutput(playerUi.GetLogView())

	go func() {
		err := discord.InitDiscordRichPresence(DiscordClientId)
		if err != nil {
			log.Println(err)
		}
	}()

	err := playerUi.StartTui()
	if err != nil {
		panic(err)
	}
}
