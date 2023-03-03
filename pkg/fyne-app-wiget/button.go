package fyneappwiget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type ButtonWithData struct {
	widget.Button
	AlternativeText string
	bind            binding.String
}

func NewButtonWithData(Data binding.String, AlternativeText string, tapped func()) *ButtonWithData {
	ButtonWithData := &ButtonWithData{}
	ButtonWithData.AlternativeText = AlternativeText
	text, err := Data.Get()
	if err != nil {
		text = AlternativeText
	}
	ButtonWithData.Text = text
	ButtonWithData.OnTapped = tapped
	ButtonWithData.bind = Data
	ButtonWithData.ExtendBaseWidget(ButtonWithData)
	ButtonWithData.Bind()
	return ButtonWithData
}

func (bwd *ButtonWithData) Bind() {
	lisiner := binding.NewDataListener(bwd.updateFromData)
	bwd.bind.AddListener(lisiner)
}

func (bwd *ButtonWithData) updateFromData() {
	if bwd.bind == nil {
		return
	}
	textSource, ok := bwd.bind.(binding.String)
	if !ok {
		return
	}
	val, err := textSource.Get()
	if err != nil {
		bwd.SetText(bwd.AlternativeText)
		fyne.LogError("Error getting current data value", err)
		return
	}
	bwd.SetText(val)
}
