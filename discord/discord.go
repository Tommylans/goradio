package discord

import (
	"errors"
	"github.com/hugolgst/rich-go/client"
	"github.com/tommylans/goradio/channels"
	"time"
)

var (
	ErrChannelIsNil = errors.New("channel is nil")
)

func InitDiscordRichPresence(clientId string) error {
	err := client.Login(clientId)
	if err != nil {
		return err
	}

	return nil
}

func UpdateDiscordPresence(channel channels.RadioStation) error {
	if channel == nil {
		return ErrChannelIsNil
	}

	now := time.Now()

	radioStationName := channel.GetName()

	activity := client.Activity{
		Details: radioStationName,
		Timestamps: &client.Timestamps{
			Start: &now,
		},
	}

	radioStationDiscordId := channel.GetDiscordSnowflakeId()
	if radioStationDiscordId != "" {
		activity.LargeImage = radioStationDiscordId
		activity.LargeText = radioStationName
	}

	err := client.SetActivity(activity)
	if err != nil {
		return err
	}

	return nil
}
