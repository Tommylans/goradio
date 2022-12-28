package radioplayer

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/tommylans/goradio/channels"
	"io"
	"log"
	"net/http"
	"time"
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

func (r *RadioPlayer) PlayRadioChannel(channel *channels.RadioChannel) error {
	r.CloseCurrentStreams()
	log.Println("Starting stream for: " + channel.Name)

	response, err := http.Get(channel.Url)
	if err != nil {
		log.Println("Error while getting stream:", err)
		return err
	}
	r.externalInputStream = response.Body

	var stream beep.Streamer

	stream, format, err := mp3.Decode(r.externalInputStream)
	if err != nil {
		return err
	}

	volume := &effects.Volume{
		Streamer: stream,
		Base:     2,
		Volume:   r.sessionVolume,
		Silent:   false,
	}

	r.volume = volume
	stream = volume

	r.play(stream, format)
	return nil
}

func (r *RadioPlayer) play(stream beep.Streamer, format beep.Format) {
	if !r.speakerInitialized {
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Millisecond*100))
		r.sampleRate = format.SampleRate
		r.speakerInitialized = true
	}

	log.Println("Samplerate:", format.SampleRate)
	if format.SampleRate != r.sampleRate {
		log.Printf("Using resampler to format from %d khz to %d khz", format.SampleRate, r.sampleRate)
		stream = beep.Resample(6, format.SampleRate, r.sampleRate, stream)
	}

	speaker.Play(stream)
}

func (r *RadioPlayer) CloseCurrentStreams() {
	if r.externalInputStream != nil {
		r.externalInputStream.Close()
	}
}
