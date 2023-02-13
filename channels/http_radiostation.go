package channels

import (
	"io"
	"net/http"
)

type RadioStationHttp struct {
	Name               string
	Url                string
	DiscordSnowflakeId string
}

func (r *RadioStationHttp) GetName() string {
	return r.Name
}

func (r *RadioStationHttp) GetDiscordSnowflakeId() string {
	return r.DiscordSnowflakeId
}

func (r *RadioStationHttp) OpenStream() (io.ReadCloser, error) {
	response, err := http.Get(r.Url)
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func (r *RadioStationHttp) GetLocation() string {
	return r.Url
}
