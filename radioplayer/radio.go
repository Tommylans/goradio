package radioplayer

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/tommylans/goradio/channels"
	"io"
	"log"
)

type RadioPlayer struct {
	sampleRate         beep.SampleRate
	speakerInitialized bool

	volume        *effects.Volume
	sessionVolume float64

	externalInputStream io.ReadCloser
}

func NewRadioPlayer() *RadioPlayer {
	return &RadioPlayer{}
}

func (r *RadioPlayer) PlayRadioChannel(channel channels.RadioStation) error {
	log.Println("Starting stream for: " + channel.GetName())

	radioStream, err := channel.OpenStream()
	if err != nil {
		return err
	}
	r.changeExternalStream(radioStream)

	var stream beep.Streamer
	stream, format, err := mp3.Decode(r.externalInputStream)
	if err != nil {
		return err
	}

	stream = r.attachVolumeControl(stream)

	return r.playStream(stream, format)
}

func (r *RadioPlayer) GetVolume() float64 {
	return r.sessionVolume
}
