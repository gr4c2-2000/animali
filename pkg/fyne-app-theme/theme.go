package fyneapptheme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type FyneFastThemeInterface interface {
	GetBackgroundColor() color.RGBA
	GetColorButton() color.RGBA
	GetColorHover() color.RGBA
}

type FyneFastTheme struct {
	BcColor     color.RGBA
	ButtonColor color.RGBA
	ColorHover  color.RGBA
}

func (fft FyneFastTheme) GetBackgroundColor() color.RGBA {
	return fft.BcColor
}
func (fft FyneFastTheme) GetColorButton() color.RGBA {
	return fft.ButtonColor
}
func (fft FyneFastTheme) GetColorHover() color.RGBA {
	return fft.ColorHover
}

type FyneTheme[T FyneFastThemeInterface] struct {
	colors T
}

func New[T FyneFastThemeInterface](colors T) FyneTheme[T] {
	new := FyneTheme[T]{colors: colors}
	return new
}

var _ fyne.Theme = (*FyneTheme[FyneFastTheme])(nil)

func (m FyneTheme[T]) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {

	switch name {
	case theme.ColorNameButton:
		return m.colors.GetColorButton()
	case theme.ColorNameHover:
		return m.colors.GetColorHover() //https://icolorpalette.com/color/2a5278
	case theme.ColorNameBackground:
		return m.colors.GetBackgroundColor() //https://icolorpalette.com/color/flat-blue
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (m FyneTheme[T]) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m FyneTheme[T]) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m FyneTheme[T]) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
