package radioplayer

import "github.com/gopxl/beep/speaker"

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
