package animali

import (
	"image/color"

	fyneapptheme "github.com/gr4c2-2000/animali/pkg/fyne-app-theme"
)

type FyneSimpleThemes struct {
	Themes  map[string]fyneapptheme.FyneTheme[fyneapptheme.FyneFastTheme]
	Default string
}

func (fst *FyneSimpleThemes) ThemeByName(name string) fyneapptheme.FyneTheme[fyneapptheme.FyneFastTheme] {
	if val, ok := fst.Themes[name]; ok {
		return val
	}
	return fst.Themes[fst.Default]
}

func InitFyneTheme() FyneSimpleThemes {
	fst := FyneSimpleThemes{}
	fst.Default = YELLOW
	fst.Themes = make(map[string]fyneapptheme.FyneTheme[fyneapptheme.FyneFastTheme])
	fst.Themes[YELLOW] = fyneapptheme.New(fyneapptheme.FyneFastTheme{
		ColorHover:  color.RGBA{244, 194, 0, 255},
		ButtonColor: color.RGBA{244, 194, 0, 255},
		BcColor:     color.RGBA{255, 214, 53, 255}})

	fst.Themes[PURPLE] = fyneapptheme.New(fyneapptheme.FyneFastTheme{
		ColorHover:  color.RGBA{109, 37, 103, 255},
		ButtonColor: color.RGBA{109, 37, 103, 255},
		BcColor:     color.RGBA{157, 54, 149, 255}})

	fst.Themes[PINK] = fyneapptheme.New(fyneapptheme.FyneFastTheme{
		ColorHover:  color.RGBA{241, 115, 138, 255},
		ButtonColor: color.RGBA{241, 115, 138, 255},
		BcColor:     color.RGBA{255, 192, 203, 255}})

	fst.Themes[BLUE] = fyneapptheme.New(fyneapptheme.FyneFastTheme{
		ColorHover:  color.RGBA{35, 170, 226, 255},
		ButtonColor: color.RGBA{35, 170, 226, 255},
		BcColor:     color.RGBA{77, 198, 226, 255}})

	fst.Themes[GREEN] = fyneapptheme.New(fyneapptheme.FyneFastTheme{
		ColorHover:  color.RGBA{0, 73, 0, 255},
		ButtonColor: color.RGBA{0, 73, 0, 255},
		BcColor:     color.RGBA{0, 114, 0, 255}})

	fst.Themes[RED] = fyneapptheme.New(fyneapptheme.FyneFastTheme{
		ColorHover:  color.RGBA{89, 0, 0, 255},
		ButtonColor: color.RGBA{89, 0, 0, 255},
		BcColor:     color.RGBA{171, 0, 0, 255}})

	return fst

}
