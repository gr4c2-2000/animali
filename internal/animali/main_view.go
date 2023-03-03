package animali

import (
	"image/color"

	appwiget "github.com/gr4c2-2000/animali/pkg/fyne-app-wiget"

	eventworker "github.com/gr4c2-2000/animali/pkg/event-worker"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type MainView struct {
	container *fyne.Container
	Img       fyne.Resource
}

func BuildMainView() *MainView {

	image := canvas.NewImageFromResource(resourceAnimaliPng)
	image.FillMode = canvas.ImageFillContain
	butMusic := appwiget.NewButtonWithData(LanguagePack.Get(MUSIC), MUSIC, func() { Queue <- eventworker.NewEvent(VIEW, MUSIC) })
	butAnimals := appwiget.NewButtonWithData(LanguagePack.Get(ANIMALS), ANIMALS, func() { Queue <- eventworker.NewEvent(VIEW, ANIMALS) })
	butThemes := appwiget.NewButtonWithData(LanguagePack.Get("App Color"), "App Color", func() { Queue <- eventworker.NewEvent(VIEW, THEME) })
	menu := container.New(layout.NewVBoxLayout(), butMusic, butAnimals, butThemes)
	empty := canvas.NewText("", color.White)
	centerGrid := container.New(layout.NewGridLayout(3), empty, menu, empty)
	Container := container.New(layout.NewGridLayout(1), image, centerGrid, empty)
	mv := MainView{container: Container}
	return &mv
}
func (m *MainView) Container() *fyne.Container {
	return m.container
}
