package channels

import (
	"context"
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

func (r *RadioStationHttp) OpenStream(ctx context.Context) (io.ReadCloser, error) {
	response, err := http.Get(r.Url) // TODO: Add a context to this so we can cancel the request
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func (r *RadioStationHttp) GetLocation() string {
	return r.Url
}
