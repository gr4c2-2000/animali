package animali

import (
	eventworker "Animali/pkg/event-worker"
	fyneappsettings "Animali/pkg/fyne-app-settings"
	fynelanguage "Animali/pkg/fyne-language"
	"context"
	"errors"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var Queue = make(chan eventworker.Event)
var LanguagePack fynelanguage.LanguagePack

type State struct {
	CurrentConteiner *fyne.Container
}

type App struct {
	State                State
	Player               *Player
	MusicScreenResources []fyne.Resource
	Themes               FyneSimpleThemes
	FyneApp              fyne.App
	Screen               map[string]Screen
	main                 fyne.Window
	Settings             Settings
	EventWoker           *eventworker.EventWoker
}
type Settings struct {
	Language  string
	ThemeName string
}

type Screen struct {
	title     string
	Conteiner *fyne.Container
}

func InitApp() *App {
	a := App{}
	a.Player = InitPayer()
	a.EventWoker = &eventworker.EventWoker{Queue: Queue}
	a.EventWoker.AddListiner(a.ThemeListiner, a.ViewListiner)
	a.EventWoker.CommandWorker()
	a.Themes = InitFyneTheme()
	LanguagePack = *fynelanguage.InitLanguagePack()
	return &a
}

func (a *App) Run() {
	a.FyneApp = app.NewWithID("test.example.com")
	fyneappsettings.InitFyneAppSettings(&a.Settings, a.FyneApp).Watch(context.TODO(), 1*time.Second)
	Queue <- eventworker.NewEvent(THEME, a.Settings.ThemeName)
	mav := BuildMainView()
	a.Screen = make(map[string]Screen)
	MainScr := Screen{title: MAIN, Conteiner: mav.container}
	a.Screen[MAIN] = MainScr
	muv := BuildMusicView(a.Player)
	Music := Screen{title: MUSIC, Conteiner: muv.Container()}
	a.Screen[MUSIC] = Music
	a.FyneApp.Lifecycle().SetOnEnteredForeground(a.Player.Stop)
	a.FyneApp.Lifecycle().SetOnExitedForeground(a.Player.Stop)
	a.main = a.FyneApp.NewWindow(TITLE)
	Queue <- eventworker.NewEvent(VIEW, MAIN)
	a.Main().ShowAndRun()
}

func (a *App) Main() fyne.Window {
	return a.main
}

func (a *App) ViewListiner(typ string, value string) {
	if typ == VIEW {
		if screen, ok := a.Screen[value]; ok {
			a.main.SetContent(screen.Conteiner)
		} else {
			fyne.LogError("", errors.New("incorect ScreenName"))
		}
	}
}

func (a *App) ThemeListiner(typ string, value string) {
	if typ == THEME {
		a.FyneApp.Settings().SetTheme(a.Themes.ThemeByName(value))
	}
}
