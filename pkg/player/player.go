package player

import (
	"bytes"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

type Player struct {
	OtoCtx *oto.Context
}

func InitPayer() (*Player, error) {
	pl := Player{}
	samplingRate := 44100
	numOfChannels := 2
	audioBitDepth := 2
	OtoCtx, readyChan, err := oto.NewContext(samplingRate, numOfChannels, audioBitDepth)
	if err != nil {
		return nil, err
	}
	pl.OtoCtx = OtoCtx
	<-readyChan
	return &pl, nil
}
func (pl *Player) Play(stop <-chan struct{}, audio []byte) (<-chan struct{}, error) {

	fileBytesReader := bytes.NewReader(audio)
	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		return nil, err
	}

	player := pl.OtoCtx.NewPlayer(decodedMp3)
	player.Play()
	end := make(chan struct{}, 0)
	go func() {
		for player.IsPlaying() {
			time.Sleep(time.Millisecond)
		}
		err = player.Close()
		if err != nil {
			panic("player.Close failed: " + err.Error())
		}
		end <- struct{}{}
		return
	}()
	go func() {
		for {
			select {
			case <-stop:
				err = player.Close()
				if err != nil {
					panic("player.Close failed: " + err.Error())
				}
				end <- struct{}{}
				return
			}
		}
	}()
	return end, nil
}
