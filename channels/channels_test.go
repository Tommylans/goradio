package channels

import (
	"context"
	"testing"
)

func TestChannelsConfiguration(t *testing.T) {
	for _, channel := range RadioChannels {
		t.Run(channel.GetName(), func(t *testing.T) {
			if channel.GetName() == "" {
				t.Errorf("Channel name is required")
			}

			stream, err := channel.OpenStream(context.Background())
			if err != nil {
				t.Errorf("Channel returned an error: %s", err)
			}

			stream.Close()
		})
	}
}
