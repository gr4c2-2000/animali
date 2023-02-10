package animali

import (
	fyneappsettings "Animali/pkg/fyne-app-settings"
	fynelanguage "Animali/pkg/fyne-language"
	"context"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var ContentChannal = make(chan string)
var LanguagePack fynelanguage.LanguagePack

type State struct {
	CurrentConteiner *fyne.Container
}

type App struct {
	State                State
	Player               *Player
	MusicScreenResources []fyne.Resource
	Theme                yellowTheme
	FyneApp              fyne.App
	Screen               map[string]Screen
	main                 fyne.Window
	Settings             Settings
}
type Settings struct {
	Language string
	ThemeID  int
}

type Screen struct {
	title     string
	Conteiner *fyne.Container
}

func InitApp() *App {
	a := App{}
	a.Player = InitPayer()
	a.SetContentWorker()

	LanguagePack = *fynelanguage.InitLanguagePack()
	return &a
}

func (a *App) Run() {
	a.FyneApp = app.NewWithID("test.example.com")
	fyneappsettings.InitFyneAppSettings(&a.Settings, a.FyneApp).Listiner(context.TODO(), 1*time.Second)

	a.FyneApp.Settings().SetTheme(&yellowTheme{})
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
	ContentChannal <- MAIN
	a.Main().ShowAndRun()
}

func (a *App) Main() fyne.Window {
	return a.main
}

func (a *App) SetContentWorker() {
	go func() {
		for val := range ContentChannal {
			if screen, ok := a.Screen[val]; ok {
				a.main.SetContent(screen.Conteiner)
			}
		}
	}()
}
