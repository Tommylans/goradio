package discord

import (
	"errors"
	"github.com/hugolgst/rich-go/client"
	"radio/channels"
	"time"
)

func InitDiscordRichPresence(clientId string) error {
	err := client.Login(clientId)
	if err != nil {
		return err
	}

	return nil
}

func UpdateDiscordPresence(channel *channels.RadioChannel) error {
	if channel == nil {
		return errors.New("channel is nil")
	}

	now := time.Now()

	activity := client.Activity{
		Details: channel.Name,
		Timestamps: &client.Timestamps{
			Start: &now,
		},
	}

	if channel.DiscordSnowflakeId != "" {
		activity.LargeImage = channel.DiscordSnowflakeId
		activity.LargeText = channel.Name
	}

	err := client.SetActivity(activity)
	if err != nil {
		return err
	}

	return nil
}
