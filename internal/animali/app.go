package animali

import (
	"context"
	"errors"
	"os"
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
	screenMap            map[string]Screen
	main                 fyne.Window
	Settings             Settings
	EventWoker           *eventworker.EventWoker
}

func (a *App) AddScreen(name string, conteiner *fyne.Container) {
	screen := Screen{title: name, Conteiner: conteiner}
	a.screenMap[name] = screen
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
	ctx, cancel := context.WithCancel(context.Background())
	a.dependencies(ctx)
	a.fyneSetup(ctx, cancel)
	a.screen()
	a.main = a.FyneApp.NewWindow(TITLE)
	return &a
}
func (a *App) Run() {
	Queue <- eventworker.NewEvent(THEME, a.Settings.ThemeName)
	Queue <- eventworker.NewEvent(VIEW, MAIN)
	a.main.ShowAndRun()
}

func (a *App) fyneSetup(ctx context.Context, cancel context.CancelFunc) {
	a.FyneApp = app.NewWithID("github.com/gr4c2-2000/animali")
	a.FyneApp.Lifecycle().SetOnEnteredForeground(a.Player.Stop)
	a.FyneApp.Lifecycle().SetOnExitedForeground(a.Player.Stop)
	a.FyneApp.Lifecycle().SetOnStopped(cancel)
	fyneappsettings.InitBridge(&a.Settings, a.FyneApp.Preferences()).Watch(context.TODO(), 1*time.Second)
}
func (a *App) dependencies(ctx context.Context) {
	err, Player := InitPayer()
	handleInitError(err)
	a.EventWoker = eventworker.New(ctx, Queue, a.ThemeListiner, a.ViewListiner)
	err = a.EventWoker.Worker()
	handleInitError(err)
	a.Player = Player
	a.Themes = InitFyneTheme()
	LanguagePack = *fynelanguage.InitLanguagePack()
}
func (a *App) screen() {
	a.screenMap = make(map[string]Screen)
	a.AddScreen(MAIN, MainView())
	a.AddScreen(MUSIC, SubView(MusicView(a.Player)))
	a.AddScreen(THEME, SubView(ThemeView(a.Themes)))
}

func (a *App) ViewListiner(typ string, value string) {
	if typ == VIEW {
		if screen, ok := a.screenMap[value]; ok {
			a.main.SetContent(screen.Conteiner)
			a.Player.Stop()
		} else {
			fyne.LogError("", errors.New("incorect ScreenName"))
		}
	}
}

func (a *App) ThemeListiner(typ string, value string) {
	if typ == THEME {
		a.FyneApp.Settings().SetTheme(a.Themes.ThemeByName(value))
		a.Settings.ThemeName = value
		for _, val := range a.screenMap {
			val.Conteiner.Refresh()
		}
	}
}

func handleInitError(err error) {
	if err != nil {
		fyne.LogError("", err)
		os.Exit(0)
	}
}
