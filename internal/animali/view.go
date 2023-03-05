package animali

import (
	"image/color"

	eventworker "github.com/gr4c2-2000/animali/pkg/event-worker"
	fyneappview "github.com/gr4c2-2000/animali/pkg/fyne-app-view"
	appwiget "github.com/gr4c2-2000/animali/pkg/fyne-app-wiget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type SoundGridView struct {
	fyneappview.GridView
	Field []MusicViewItem
}
type MusicViewItem struct {
	Button fyne.Resource
	Audio  fyne.Resource
}

func soundGridView(Player *Player, mvi []MusicViewItem) *SoundGridView {
	items := make([]fyneappview.ImageGridField, 0)
	for _, v := range mvi {

		Player.AddToPlaylist(v.Audio, v.Audio.Name())
		item := fyneappview.ImageGridField{Img: v.Button}
		name := v.Audio.Name()
		function := func() {
			Player.SetSong(name)
			Player.Play()
		}
		item.OnTapped = function
		items = append(items, item)
	}
	gridItems := fyneappview.ImageGridFields{Fields: items}
	GridView := fyneappview.NewGridView(3, gridItems)
	return &SoundGridView{*GridView, mvi}
}

func MusicView(Player *Player) *fyne.Container {

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
	return soundGridView(Player, mvi).Container()
}

func AnimalView(Player *Player) *fyne.Container {
	// TODO: add create resources, need to found sounds and images
	mvi := make([]MusicViewItem, 0)

	return soundGridView(Player, mvi).Container()
}

func MainView() *fyne.Container {
	image := canvas.NewImageFromResource(resourceAnimaliPng)
	image.FillMode = canvas.ImageFillContain
	butMusic := appwiget.NewButtonWithData(LanguagePack.Get(MUSIC), MUSIC, func() { Queue <- eventworker.NewEvent(VIEW, MUSIC) })
	butAnimals := appwiget.NewButtonWithData(LanguagePack.Get(ANIMALS), ANIMALS, func() { Queue <- eventworker.NewEvent(VIEW, ANIMALS) })
	butThemes := appwiget.NewButtonWithData(LanguagePack.Get("App Color"), "App Color", func() { Queue <- eventworker.NewEvent(VIEW, THEME) })
	menu := container.New(layout.NewVBoxLayout(), butMusic, butAnimals, butThemes)
	empty := canvas.NewText("", color.White)
	centerGrid := container.New(layout.NewGridLayout(3), empty, menu, empty)
	Container := container.New(layout.NewGridLayout(1), image, centerGrid, empty)
	return Container
}

func SubView(me *fyne.Container) *fyne.Container {
	return container.NewBorder(TopBar(), nil, nil, nil, me)
}

func ThemeView(fst FyneSimpleThemes) *fyne.Container {
	items := make([]fyneappview.ColorGridField, 0)
	for name, v := range fst.Themes {

		item := fyneappview.ColorGridField{Color: v.Colors.BcColor}
		color := name
		function := func() {
			Queue <- eventworker.NewEvent(THEME, color)
		}
		item.OnTapped = function
		items = append(items, item)
	}
	gridItems := fyneappview.ColorGridFields{Fields: items}
	GridView := fyneappview.NewGridView(3, gridItems)
	return GridView.Container()
}

func TopBar() fyne.CanvasObject {
	function := func() {
		Queue <- eventworker.NewEvent(VIEW, MAIN)
	}
	backbt := widget.NewButton("BACK", function)
	return backbt
}
