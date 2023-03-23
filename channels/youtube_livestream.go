package channels

import (
	"bytes"
	"context"
	"errors"
	"github.com/Eyevinn/mp4ff/aac"
	"github.com/Eyevinn/mp4ff/mp4"
	"github.com/kkdai/youtube/v2"
	"io"
	"log"
	"net/http"
)

type YoutubeLivestream struct {
	Name string
	Url  string
}

func (y *YoutubeLivestream) GetName() string {
	return y.Name
}

func (y *YoutubeLivestream) GetDiscordSnowflakeId() string {
	return "youtube"
}

type Track struct {
	trackID   uint32
	hdlrType  string
	timeScale uint64
	trak      *mp4.TrakBox
	trex      *mp4.TrexBox
	samples   []mp4.FullSample
}

func getTracksAndSamplesFromMultiTrackFragmentedFile(ifd io.Reader) (tracks []*Track, err error) {
	parsedMp4, err := mp4.DecodeFile(ifd)
	if err != nil {
		log.Fatalln(err)
	}
	traks := parsedMp4.Moov.Traks

	for _, trak := range traks {
		track := &Track{}
		track.trak = trak
		track.trackID = trak.Tkhd.TrackID
		track.timeScale = uint64(trak.Mdia.Mdhd.Timescale)
		track.hdlrType = trak.Mdia.Hdlr.HandlerType
		for _, trex := range parsedMp4.Moov.Mvex.Trexs {
			if trex.TrackID == track.trackID {
				track.trex = trex
				break
			}
		}
		tracks = append(tracks, track)
	}

	for _, seg := range parsedMp4.Segments {
		for _, frag := range seg.Fragments {
			for _, track := range tracks {
				samples, err := frag.GetFullSamples(track.trex)
				if err != nil {
					log.Fatalln(err)
				}
				track.samples = append(track.samples, samples...)
			}
		}
	}
	return tracks, nil
}

func (y *YoutubeLivestream) OpenStream(ctx context.Context) (io.ReadCloser, error) {
	client := youtube.Client{}

	video, err := client.GetVideo(y.Url)
	if err != nil {
		return nil, err
	}

	channels := video.Formats.WithAudioChannels()

	channels.Sort()

	if len(channels) == 0 {
		return nil, errors.New("no youtube stream found")
	}

	format := channels[0]

	streamUrl, err := client.GetStreamURLContext(ctx, video, &format)
	if err != nil {
		return nil, err
	}

	// TODO: Add implementation for seperated audio stream
	resp, err := http.Get(streamUrl)
	if err != nil {
		return nil, err
	}

	//tracks, err := getTracksAndSamplesFromMultiTrackFragmentedFile(resp.Body)
	//if err != nil {
	//	return nil, err
	//}
	//
	//reader := bytes.NewReader(tracks[0].samples[0].Data)
	//config, err := aac.DecodeAudioSpecificConfig(reader)
	//if err != nil {
	//	panic(err)
	//}

	file, err := mp4.DecodeFile(resp.Body)
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(file.LastSegment().LastFragment().Mdat.Data)
	config, err := aac.DecodeAudioSpecificConfig(reader)
	if err != nil {
		panic(err)
	}

	print(config)

	panic("not implemented")
}

func (y *YoutubeLivestream) GetLocation() string {
	return y.Url
}
