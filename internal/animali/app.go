package animali

import (
	fyneappsettings "Animali/pkg/fyne-app-settings"
	fynelanguage "Animali/pkg/fyne-language"
	"context"
	"errors"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var ExecCommand = make(chan Command)
var LanguagePack fynelanguage.LanguagePack

type Command struct {
	Type  string
	Value string
}
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
	a.CommandWorker()
	a.Themes = InitFyneTheme()
	LanguagePack = *fynelanguage.InitLanguagePack()
	return &a
}

func (a *App) Run() {
	a.FyneApp = app.NewWithID("test.example.com")
	fyneappsettings.InitFyneAppSettings(&a.Settings, a.FyneApp).Listiner(context.TODO(), 1*time.Second)
	ExecCommand <- Command{THEME, a.Settings.ThemeName}
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
	ExecCommand <- Command{VIEW, MAIN}
	a.Main().ShowAndRun()
}

func (a *App) Main() fyne.Window {
	return a.main
}

func (a *App) CommandWorker() {

	go func() {
		for Command := range ExecCommand {
			switch Command.Type {
			case VIEW:
				if screen, ok := a.Screen[Command.Value]; ok {
					a.main.SetContent(screen.Conteiner)
				} else {
					fyne.LogError("", errors.New("incorect ScreenName"))
				}
			case THEME:
				a.FyneApp.Settings().SetTheme(a.Themes.ThemeByName(Command.Value))
			default:
				fyne.LogError("", errors.New("Incorect Command"))
			}
		}
	}()
}
