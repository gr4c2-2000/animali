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
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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
	err := a.EventWoker.Worker()
	if err != nil {
		fyne.LogError("", err)
		panic(err)
	}
	a.Themes = InitFyneTheme()
	LanguagePack = *fynelanguage.InitLanguagePack()
	a.FyneApp = app.NewWithID("test.example.com")

	a.Screen = make(map[string]Screen)
	fyneappsettings.InitBridge(&a.Settings, a.FyneApp.Preferences()).Watch(context.TODO(), 1*time.Second)
	Queue <- eventworker.NewEvent(THEME, a.Settings.ThemeName)
	mav := BuildMainView()
	muv := BuildMusicView(a.Player)
	tv := BuildThemeView(a.Themes)
	MainScr := Screen{title: MAIN, Conteiner: mav.container}
	function := func() {
		Queue <- eventworker.NewEvent(VIEW, MAIN)
	}
	backbt := widget.NewButton("BACK", function)
	Music2 := container.NewBorder(backbt, nil, nil, nil, muv.Container())
	Music := Screen{title: MUSIC, Conteiner: Music2}

	Theme := Screen{title: THEME, Conteiner: container.NewBorder(backbt, nil, nil, nil, tv.Container())}
	a.Screen[MAIN] = MainScr
	a.Screen[MUSIC] = Music
	a.Screen[THEME] = Theme
	a.FyneApp.Lifecycle().SetOnEnteredForeground(a.Player.Stop)
	a.FyneApp.Lifecycle().SetOnExitedForeground(a.Player.Stop)
	a.main = a.FyneApp.NewWindow(TITLE)

	return &a
}

func (a *App) Run() {

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
