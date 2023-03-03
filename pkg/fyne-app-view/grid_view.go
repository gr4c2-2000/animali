package fyneappview

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type GridFields interface {
	GetFields() []fyne.CanvasObject
}
type GridView struct {
	Fields    GridFields
	GridSize  int
	items     []fyne.CanvasObject
	container *fyne.Container
}

func NewGridView(size int, gf GridFields) *GridView {
	grid := GridView{GridSize: size, Fields: gf}
	grid.items = gf.GetFields()
	return &grid
}
func (igw *GridView) Container() *fyne.Container {
	igw.container = container.New(layout.NewGridLayout(igw.GridSize), igw.items...)
	return igw.container
}
