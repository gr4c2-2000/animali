package fyneappview

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ColorGridFields struct {
	Fields []ColorGridField
	Frame  fyne.Resource
}

type ColorGridField struct {
	Color    color.RGBA
	OnTapped func()
}

func (igf ColorGridFields) GetFields() []fyne.CanvasObject {
	slc := make([]fyne.CanvasObject, 0)
	for _, field := range igf.Fields {
		image := canvas.NewImageFromResource(igf.Frame)
		image.FillMode = canvas.ImageFillContain
		imBG := canvas.NewImageFromImage(NewImage(field.Color, 5))
		imBG.FillMode = canvas.ImageFillContain
		b := widget.NewButton("", field.OnTapped)
		box := container.NewPadded(b, imBG, image)
		slc = append(slc, box)
	}
	return slc
}

// TODO: Refactior
func NewImage(col color.RGBA, margin int) *image.RGBA {
	width := 100
	height := 100
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if x > margin && x < (width-margin) && y > margin && y < (height-margin) {
				img.Set(x, y, col)
			}
		}
	}
	return img
}
