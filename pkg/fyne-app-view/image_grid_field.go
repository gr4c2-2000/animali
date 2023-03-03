package fyneappview

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ImageGridFields struct {
	Fields []ImageGridField
}

type ImageGridField struct {
	Img      fyne.Resource
	OnTapped func()
	label    string
}

func (igf ImageGridFields) GetFields() []fyne.CanvasObject {
	slc := make([]fyne.CanvasObject, 0)
	for _, field := range igf.Fields {
		image := canvas.NewImageFromResource(field.Img)
		image.FillMode = canvas.ImageFillContain
		b := widget.NewButton(field.label, field.OnTapped)
		box := container.NewPadded(b, image)
		slc = append(slc, box)
	}
	return slc
}
