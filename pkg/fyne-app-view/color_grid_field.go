package fyneappview

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ColorGridFields struct {
	Fields []ColorGridField
}

type ColorGridField struct {
	Color    color.RGBA
	OnTapped func()
}

func (igf ColorGridFields) GetFields() []fyne.CanvasObject {
	slc := make([]fyne.CanvasObject, 0)
	for _, field := range igf.Fields {
		b := widget.NewButton("", field.OnTapped)
		rectangle := canvas.NewRectangle(field.Color)
		box := container.NewPadded(b, rectangle)
		slc = append(slc, box)
	}
	return slc
}
