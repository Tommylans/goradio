package player

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"io"
	"log"
	"net/http"
	"radio/channels"
	"time"
)

type RadioPlayer struct {
	sampleRate beep.SampleRate
	logger     *log.Logger

	externalInputStream io.ReadCloser
}

func (r *RadioPlayer) SetLogger(logger *log.Logger) {
	r.logger = logger
}

func NewRadioPlayer() *RadioPlayer {
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Millisecond*100))

	return &RadioPlayer{sampleRate: sr}
}

func (r *RadioPlayer) Play() {
	speaker.Clear()
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

	decode, format, err := mp3.Decode(r.externalInputStream)
	if err != nil {
		return err
	}

	var stream beep.Streamer = decode
	r.logger.Println("Samplerate:", format.SampleRate)
	if format.SampleRate != r.sampleRate {
		r.logger.Printf("Using resampler to format from %d khz to %d khz", int(format.SampleRate), int(r.sampleRate))
		stream = beep.Resample(6, format.SampleRate, r.sampleRate, stream)
	}

	speaker.Play(stream)
	return nil
}

func (r *RadioPlayer) Close() {
	if r.externalInputStream != nil {
		r.externalInputStream.Close()
	}
}
