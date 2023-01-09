package animali

import (
	fyneappview "Animali/pkg/fyne-app-view"

	"fyne.io/fyne/v2"
)

type MusicView struct {
	fyneappview.ImageGridView
}

func Build(Player *Player) *MusicView {
	MusicScreenResources := []fyne.Resource{resourceButton1Png, resourceButton1Png, resourceButton3Png, resourceButton1Png,
		resourceButton3Png, resourceButton1Png, resourceButton2Png, resourceButton3Png, resourceButton3Png}

	items := make([]fyneappview.ImageGridField, 0)
	for i, res := range MusicScreenResources {
		item := fyneappview.ImageGridField{Img: res}
		var function func()
		if i%2 == 0 {
			function = func() {
				Player.SetSong(resourceSongMp3.StaticName)
				Player.Play()
			}
		} else {
			function = func() {
				Player.SetSong(resourceShortMp3.StaticName)
				Player.Play()
			}
		}
		item.OnTapped = function
		items = append(items, item)
	}
	GridView := fyneappview.NewImageGridView(3, items...)
	return &MusicView{*GridView}

}
