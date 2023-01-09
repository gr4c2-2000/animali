package fyneappview

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type ImageGridView struct {
	Fields    []ImageGridField
	GridSize  int
	items     []fyne.CanvasObject
	container *fyne.Container
}
type ImageGridField struct {
	Img      fyne.Resource
	OnTapped func()
	label    string
}

func NewImageGridView(size int, slc ...ImageGridField) *ImageGridView {
	grid := ImageGridView{GridSize: size}
	for _, it := range slc {
		grid.Fields = append(grid.Fields, it)
	}
	grid.MakeItems().MakeContainer()
	return &grid
}
func NewImageGridViewConteiner(size int, slc ...ImageGridField) *fyne.Container {
	grid := ImageGridView{GridSize: size}
	for _, it := range slc {
		grid.Fields = append(grid.Fields, it)
	}
	grid.MakeItems().MakeContainer()
	return grid.container
}
func (igw *ImageGridView) Container() *fyne.Container {
	return igw.container
}
func (igw *ImageGridView) MakeContainer() *ImageGridView {
	igw.container = container.New(layout.NewGridLayout(igw.GridSize), igw.items...)
	return igw
}
func (igw *ImageGridView) MakeItems() *ImageGridView {
	slc := make([]fyne.CanvasObject, 0)
	for _, field := range igw.Fields {
		image := canvas.NewImageFromResource(field.Img)
		image.FillMode = canvas.ImageFillContain
		b := widget.NewButton(field.label, field.OnTapped)
		box := container.NewPadded(b, image)
		slc = append(slc, box)
	}
	igw.items = slc
	return igw
}
