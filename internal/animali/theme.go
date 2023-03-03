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
		ColorHover:  color.RGBA{244, 194, 0, 255},
		ButtonColor: color.RGBA{244, 194, 0, 255},
		BcColor:     color.RGBA{157, 54, 149, 255}})

	fst.Themes[PINK] = fyneapptheme.New(fyneapptheme.FyneFastTheme{
		ColorHover:  color.RGBA{244, 194, 0, 255},
		ButtonColor: color.RGBA{244, 194, 0, 255},
		BcColor:     color.RGBA{255, 192, 203, 255}})

	fst.Themes[BLUE] = fyneapptheme.New(fyneapptheme.FyneFastTheme{
		ColorHover:  color.RGBA{88, 198, 226, 255},
		ButtonColor: color.RGBA{88, 198, 226, 255},
		BcColor:     color.RGBA{77, 198, 226, 255}})

	fst.Themes[GREEN] = fyneapptheme.New(fyneapptheme.FyneFastTheme{
		ColorHover:  color.RGBA{244, 194, 0, 255},
		ButtonColor: color.RGBA{244, 194, 0, 255},
		BcColor:     color.RGBA{51, 255, 51, 255}})

	fst.Themes[RED] = fyneapptheme.New(fyneapptheme.FyneFastTheme{
		ColorHover:  color.RGBA{244, 194, 0, 255},
		ButtonColor: color.RGBA{244, 194, 0, 255},
		BcColor:     color.RGBA{51, 255, 51, 255}})

	return fst

}
