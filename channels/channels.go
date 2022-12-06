package channels

type RadioChannel struct {
	Name        string
	Url         string
	SnowflakeId string
}

var (
	RadioChannels = []RadioChannel{
		{Name: "SlamFM", Url: "https://stream.slam.nl/slam_mp3", SnowflakeId: "slamfm"},
		{Name: "Veronica", Url: "https://playerservices.streamtheworld.com/api/livestream-redirect/VERONICA.mp3"},
		{Name: "NPO Radio 1", Url: "https://icecast.omroep.nl/radio1-bb-mp3"},
		{Name: "NPO Radio 2", Url: "https://icecast.omroep.nl/radio2-bb-mp3", SnowflakeId: "nporadio2"},
		{Name: "NPO 3FM", Url: "https://icecast.omroep.nl/3fm-bb-mp3", SnowflakeId: "3fm"},
		{Name: "NPO FunX", Url: "https://icecast.omroep.nl/funx-bb-mp3"},
	}
)
