package radioplayer

import "github.com/faiface/beep/speaker"

func (r *RadioPlayer) Stop() {
	speaker.Clear()
}

func (r *RadioPlayer) Mute() {
	if r.volume != nil {
		speaker.Lock()
		r.volume.Silent = !r.volume.Silent
		speaker.Unlock()
	}
}

func (r *RadioPlayer) IncreaseVolume() {
	r.changeVolume(0.5)
}

func (r *RadioPlayer) ResetVolume() {
	r.changeVolume(-r.sessionVolume)
}

func (r *RadioPlayer) DecreaseVolume() {
	r.changeVolume(-0.5)
}

func (r *RadioPlayer) changeVolume(change float64) {
	if r.volume != nil {
		speaker.Lock()
		r.volume.Volume += change
		r.sessionVolume = r.volume.Volume
		speaker.Unlock()
	}
}
