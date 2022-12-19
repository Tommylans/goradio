package channels

type RadioChannel struct {
	Name               string
	Url                string
	DiscordSnowflakeId string
}

var (
	RadioChannels = []*RadioChannel{
		{Name: "SlamFM", Url: "https://stream.slam.nl/slam_mp3", DiscordSnowflakeId: "slamfm"},
		{Name: "Veronica", Url: "https://playerservices.streamtheworld.com/api/livestream-redirect/VERONICA.mp3"},
		{Name: "SkyRadio", Url: "https://playerservices.streamtheworld.com/api/livestream-redirect/SKYRADIO.mp3", DiscordSnowflakeId: "skyradio"},
		{Name: "NPO Radio 1", Url: "https://icecast.omroep.nl/radio1-bb-mp3"},
		{Name: "NPO Radio 2", Url: "https://icecast.omroep.nl/radio2-bb-mp3", DiscordSnowflakeId: "nporadio2"},
		{Name: "NPO Radio 2 Soul & Jazz", Url: "https://icecast.omroep.nl/radio6-bb-mp3", DiscordSnowflakeId: "nporadio2"},
		{Name: "NPO 3FM", Url: "https://icecast.omroep.nl/3fm-bb-mp3", DiscordSnowflakeId: "3fm"},
		{Name: "NPO FunX", Url: "https://icecast.omroep.nl/funx-bb-mp3"},
		{Name: "538", Url: "https://playerservices.streamtheworld.com/api/livestream-redirect/RADIO538.mp3", DiscordSnowflakeId: "538"},
		{Name: "538 Verrückte Stunden", Url: "https://playerservices.streamtheworld.com/api/livestream-redirect/TLPSTR21.mp3", DiscordSnowflakeId: "538"},
	}
)
