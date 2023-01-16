package animali

import (
	fyneappview "Animali/pkg/fyne-app-view"
	"fmt"

	"fyne.io/fyne/v2"
)

type SoundGridView struct {
	fyneappview.ImageGridView
	Field []MusicViewItem
}
type MusicViewItem struct {
	Button fyne.Resource
	Audio  fyne.Resource
}

func BuildSoundGridView(Player *Player, mvi []MusicViewItem) *SoundGridView {
	items := make([]fyneappview.ImageGridField, 0)
	for _, v := range mvi {

		Player.AddToPlaylist(v.Audio, v.Audio.Name())
		fmt.Println(v.Audio.Name())
		item := fyneappview.ImageGridField{Img: v.Button}
		name := v.Audio.Name()
		function := func() {
			Player.SetSong(name)
			Player.Play()
		}
		item.OnTapped = function
		items = append(items, item)
	}
	GridView := fyneappview.NewImageGridView(3, items...)
	return &SoundGridView{*GridView, mvi}
}

func BuildMusicView(Player *Player) *SoundGridView {

	mvi := make([]MusicViewItem, 0)
	mvi = append(mvi, MusicViewItem{resourceButton1Png, resourceSong1Mp3})
	mvi = append(mvi, MusicViewItem{resourceButton2Png, resourceSong2Mp3})
	mvi = append(mvi, MusicViewItem{resourceButton3Png, resourceSong3Mp3})
	mvi = append(mvi, MusicViewItem{resourceButton4Png, resourceSong4Mp3})
	mvi = append(mvi, MusicViewItem{resourceButton5Png, resourceSong5Mp3})
	mvi = append(mvi, MusicViewItem{resourceButton6Png, resourceSong6Mp3})
	mvi = append(mvi, MusicViewItem{resourceButton7Png, resourceSong7Mp3})
	mvi = append(mvi, MusicViewItem{resourceButton8Png, resourceSong8Mp3})
	mvi = append(mvi, MusicViewItem{resourceButton9Png, resourceSong9Mp3})
	return BuildSoundGridView(Player, mvi)
}

func BuildAnimalView(Player *Player) *SoundGridView {
	// add create resources
	mvi := make([]MusicViewItem, 0)

	return BuildSoundGridView(Player, mvi)
}
