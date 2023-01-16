package animali

import (
	"Animali/pkg/player"
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
)

type Player struct {
	list      map[string]fyne.Resource
	current   fyne.Resource
	stop      chan struct{}
	playing   bool
	byteschan chan []byte
	Player    *player.Player
}

func InitPayer() *Player {
	ls := make(map[string]fyne.Resource)
	stop := make(chan struct{})
	pla, err := player.InitPayer()
	if err != nil {
		panic(err)
	}
	pl := Player{list: ls, stop: stop, Player: pla}
	pl.Worker()
	return &pl
}

func (p *Player) AddToPlaylist(fr fyne.Resource, name string) {
	p.list[name] = fr
	if p.current == nil {
		p.current = fr
	}
}
func (p *Player) Stop() {
	if p.playing {
		p.stop <- struct{}{}
	}
	p.playing = false
}
func (p *Player) SetSong(name string) error {
	if val, ok := p.list[name]; ok {
		p.current = val
		return nil
	}
	return errors.New("No song " + name)
}

func (p *Player) Play() {
	if p.playing {
		p.Stop()
	}
	p.byteschan <- p.current.Content()
}

func (p *Player) Worker() {
	p.byteschan = make(chan []byte)
	go func() {
		for data := range p.byteschan {
			p.playing = true
			stoped, err := p.Player.Play(p.stop, data)
			if err != nil {
				fmt.Println(err)
			}
			<-stoped
			p.playing = false
		}
	}()
}
