package radioplayer

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/tommylans/goradio/channels"
	"io"
	"log"
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

func (r *RadioPlayer) playStream(stream beep.Streamer, format beep.Format) error {
	sampleRate := format.SampleRate

	if !r.speakerInitialized {
		err := r.initializeSpeaker(sampleRate)
		if err != nil {
			return err
		}
	}

	stream = r.ensureSampleRate(stream, sampleRate)

	speaker.Play(stream)

	return nil
}

func (r *RadioPlayer) initializeSpeaker(sampleRate beep.SampleRate) error {
	err := speaker.Init(sampleRate, sampleRate.N(time.Millisecond*100))
	if err != nil {
		return err
	}

	r.sampleRate = sampleRate
	r.speakerInitialized = true

	return nil
}

func (r *RadioPlayer) changeExternalStream(stream io.ReadCloser) {
	if r.externalInputStream != nil {
		r.externalInputStream.Close()
	}

	r.externalInputStream = stream
}

func (r *RadioPlayer) ensureSampleRate(streamer beep.Streamer, targetSampleRate beep.SampleRate) beep.Streamer {
	if targetSampleRate != r.sampleRate {
		log.Printf("Using resampler to format from %d khz to %d khz", targetSampleRate, r.sampleRate)

		return beep.Resample(6, targetSampleRate, r.sampleRate, streamer)
	}

	return streamer
}

func (r *RadioPlayer) attachVolumeControl(streamer beep.Streamer) beep.Streamer {
	volume := &effects.Volume{
		Streamer: streamer,
		Base:     2,
		Volume:   r.sessionVolume,
		Silent:   false,
	}

	r.volume = volume

	return volume
}

func (r *RadioPlayer) changeVolume(change float64) {
	if r.volume != nil {
		speaker.Lock()
		r.volume.Volume += change
		r.sessionVolume = r.volume.Volume
		speaker.Unlock()
	}
}

func (r *RadioPlayer) Close() error {
	if r.externalInputStream != nil {
		return r.externalInputStream.Close()
	}

	return nil
}
