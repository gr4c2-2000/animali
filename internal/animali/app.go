package animali

import (
	"context"
	"errors"
	"time"

	fynelanguage "github.com/gr4c2-2000/animali/pkg/fyne-language"

	fyneappsettings "github.com/gr4c2-2000/animali/pkg/fyne-app-settings"

	eventworker "github.com/gr4c2-2000/animali/pkg/event-worker"

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

func (a *App) AddScreen(name string, conteiner *fyne.Container) {
	screen := Screen{title: name, Conteiner: conteiner}
	a.Screen[name] = screen
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
	err := a.EventWoker.Worker()
	if err != nil {
		fyne.LogError("", err)
		panic(err)
	}
	a.Themes = InitFyneTheme()
	LanguagePack = *fynelanguage.InitLanguagePack()
	a.FyneApp = app.NewWithID("test.example.com")
	a.Screen = make(map[string]Screen)
	a.AddScreen(MAIN, MainView())
	a.AddScreen(MUSIC, SubView(MusicView(a.Player)))
	fyneappsettings.InitBridge(&a.Settings, a.FyneApp.Preferences()).Watch(context.TODO(), 1*time.Second)
	a.AddScreen(THEME, SubView(ThemeView(a.Themes)))
	a.FyneApp.Lifecycle().SetOnEnteredForeground(a.Player.Stop)
	a.FyneApp.Lifecycle().SetOnExitedForeground(a.Player.Stop)
	a.main = a.FyneApp.NewWindow(TITLE)
	return &a
}

func (a *App) Run() {

	Queue <- eventworker.NewEvent(THEME, a.Settings.ThemeName)
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
		for _, val := range a.Screen {
			val.Conteiner.Refresh()
		}
	}
}
