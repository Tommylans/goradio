package channels

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
)

type RadioStationType string

const (
	ExternalRadioStationsSourceUrl = "https://raw.githubusercontent.com/Tommylans/goradio/master/assets/radio-channels.json"
)

const (
	RadioStationTypeHttpMp3 RadioStationType = "http-mp3"
)

type ExternalRadioStation struct {
	Type               RadioStationType `json:"type"`
	Name               string           `json:"name"`
	Url                string           `json:"url"`
	DiscordSnowflakeId string           `json:"discordSnowflakeId"`
}

func (r *ExternalRadioStation) ConstructRadioStation() (RadioStation, error) {
	switch r.Type {
	case RadioStationTypeHttpMp3:
		return &RadioStationHttp{Name: r.Name, Url: r.Url, DiscordSnowflakeId: r.DiscordSnowflakeId}, nil
	}

	return nil, errors.New("unknown radio station type " + string(r.Type))
}

func FetchExternalRadioChannels() ([]RadioStation, error) {
	response, err := http.Get(ExternalRadioStationsSourceUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status code " + strconv.Itoa(response.StatusCode))
	}

	allBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var externalRadioStations []ExternalRadioStation
	err = json.Unmarshal(allBytes, &externalRadioStations)
	if err != nil {
		return nil, err
	}

	var radioStations []RadioStation
	for _, externalRadioStation := range externalRadioStations {
		radioStation, err := externalRadioStation.ConstructRadioStation()
		if err != nil {
			log.Println("Maybe the radio channel is not compatible with this version: ", err)
			continue
		}

		radioStations = append(radioStations, radioStation)
	}

	return radioStations, nil
}
