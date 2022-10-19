package player

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"io"
	"log"
	"net/http"
	"radio/channels"
)

type RadioPlayer struct {
	sampleRate int
	player     *audio.Player
	logger     *log.Logger

	externalInputStream io.ReadCloser
}

func (r *RadioPlayer) SetLogger(logger *log.Logger) {
	r.logger = logger
}

func NewRadioPlayer() *RadioPlayer {
	sampleRate := 44100
	audio.NewContext(sampleRate)

	return &RadioPlayer{sampleRate: sampleRate}
}

func (r *RadioPlayer) Play() {
	r.player.Play()
}

func (r *RadioPlayer) PlayChannel(channel channels.RadioChannel) error {
	r.Close()
	r.logger.Println("Starting stream for: " + channel.Name)

	response, err := http.Get(channel.Url)
	if err != nil {
		r.logger.Println(err)
		return err
	}
	r.externalInputStream = response.Body

	decodedStream, err := mp3.DecodeWithSampleRate(r.sampleRate, r.externalInputStream)
	if err != nil {
		r.logger.Println(err)
		return err
	}

	player, err := audio.CurrentContext().NewPlayer(decodedStream)
	if err != nil {
		r.logger.Println(err)
		return err
	}
	r.player = player

	r.player.Play()
	return nil
}

func (r *RadioPlayer) Close() {
	if r.externalInputStream != nil {
		r.externalInputStream.Close()
	}
	if r.player != nil {
		r.player.Close()
	}
}
