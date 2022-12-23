package channels

import (
	"net/http"
	"testing"
)

func TestChannelsConfiguration(t *testing.T) {
	for _, channel := range RadioChannels {
		t.Run(channel.Name, func(t *testing.T) {
			if channel.Name == "" {
				t.Errorf("Channel name is required")
			}

			if channel.Url == "" {
				t.Errorf("Channel url is required")
			}

			res, err := http.Get(channel.Url)
			if err != nil {
				t.Errorf("Channel returned an error: %s", err)
			}

			defer res.Body.Close()

			if res.StatusCode != 200 {
				t.Errorf("Channel returned a non-200 status code: %d", res.StatusCode)
			}
		})
	}
}
