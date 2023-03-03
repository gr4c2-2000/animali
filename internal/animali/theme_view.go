package animali

import (
	eventworker "github.com/gr4c2-2000/animali/pkg/event-worker"
	fyneappview "github.com/gr4c2-2000/animali/pkg/fyne-app-view"
)

func BuildThemeView(fst FyneSimpleThemes) *fyneappview.GridView {
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
	return GridView
}
