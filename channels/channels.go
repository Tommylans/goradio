package channels

import (
	"io"
)

// TODO: Think about if we just want to return a beep.Streamer or if we want to return a io.ReadCloser

type RadioStation interface {
	GetName() string
	GetDiscordSnowflakeId() string
	OpenStream() (io.ReadCloser, error)
	GetLocation() string
}

var (
	RadioChannels = []RadioStation{
		&RadioStationHttp{Name: "SlamFM", Url: "https://stream.slam.nl/slam_mp3", DiscordSnowflakeId: "slamfm"},
		&RadioStationHttp{Name: "Veronica", Url: "https://playerservices.streamtheworld.com/api/livestream-redirect/VERONICA.mp3"},
		&RadioStationHttp{Name: "SkyRadio", Url: "https://playerservices.streamtheworld.com/api/livestream-redirect/SKYRADIO.mp3", DiscordSnowflakeId: "skyradio"},
		&RadioStationHttp{Name: "NPO Radio 1", Url: "https://icecast.omroep.nl/radio1-bb-mp3"},
		&RadioStationHttp{Name: "NPO Radio 2", Url: "https://icecast.omroep.nl/radio2-bb-mp3", DiscordSnowflakeId: "nporadio2"},
		&RadioStationHttp{Name: "NPO Radio 2 Soul & Jazz", Url: "https://icecast.omroep.nl/radio6-bb-mp3", DiscordSnowflakeId: "nporadio2"},
		&RadioStationHttp{Name: "NPO 3FM", Url: "https://icecast.omroep.nl/3fm-bb-mp3", DiscordSnowflakeId: "3fm"},
		&RadioStationHttp{Name: "NPO FunX", Url: "https://icecast.omroep.nl/funx-bb-mp3"},
		&RadioStationHttp{Name: "538", Url: "https://playerservices.streamtheworld.com/api/livestream-redirect/RADIO538.mp3", DiscordSnowflakeId: "538"},
		&RadioStationHttp{Name: "538 Verrückte Stunden", Url: "https://playerservices.streamtheworld.com/api/livestream-redirect/TLPSTR21.mp3", DiscordSnowflakeId: "538"},
	}
)
