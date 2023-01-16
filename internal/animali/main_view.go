package animali

import (
	appwiget "Animali/pkg/fyne-app-wiget"
	fynelanguage "Animali/pkg/fyne-language"
	"image/color"

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
	but := appwiget.NewButtonWithData(fynelanguage.LanguageStorage.Get(MUSIC), MUSIC, func() { ContentChannal <- MUSIC })
	but2 := appwiget.NewButtonWithData(fynelanguage.LanguageStorage.Get(ANIMALS), ANIMALS, func() { ContentChannal <- ANIMALS })
	menu := container.New(layout.NewVBoxLayout(), but, but2)
	empty := canvas.NewText("", color.White)
	centerGrid := container.New(layout.NewGridLayout(3), empty, menu, empty)
	Container := container.New(layout.NewGridLayout(1), image, centerGrid, empty)
	mv := MainView{container: Container}
	return &mv
}
func (m *MainView) Container() *fyne.Container {
	return m.container
}
