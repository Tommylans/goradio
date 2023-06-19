package channels

import (
	"io"
	"log"
)

// TODO: Think about if we just want to return a beep.Streamer or if we want to return a io.ReadCloser

type RadioStation interface {
	GetName() string
	GetDiscordSnowflakeId() string
	OpenStream() (io.ReadCloser, error)
	GetLocation() string
}

var (
	RadioChannels []RadioStation
)

func init() {
	radioChannels, err := FetchExternalRadioChannels()
	if err != nil {
		log.Fatal("Error fetching external radio channels: ", err)
	}

	RadioChannels = append(RadioChannels, radioChannels...)
}
