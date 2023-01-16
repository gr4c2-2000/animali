package animali

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var ContentChannal = make(chan string)

type State struct {
	CurrentConteiner *fyne.Container
}

type App struct {
	State                State
	Player               *Player
	MusicScreenResources []fyne.Resource
	Theme                myTheme
	FyneApp              fyne.App
	Screen               map[string]Screen
	main                 fyne.Window
}

type Screen struct {
	title     string
	Conteiner *fyne.Container
}

func InitApp() *App {
	a := App{}
	a.Player = InitPayer()
	a.SetContentWorker()
	return &a
}

func (a *App) Run() {
	a.FyneApp = app.New()
	a.FyneApp.Settings().SetTheme(&myTheme{})
	mav := BuildMainView()
	a.Screen = make(map[string]Screen)
	MainScr := Screen{title: "Main", Conteiner: mav.container}
	a.Screen["Main"] = MainScr
	muv := BuildMusicView(a.Player)
	Music := Screen{title: "Music", Conteiner: muv.Container()}
	a.Screen["Music"] = Music
	a.FyneApp.Lifecycle().SetOnEnteredForeground(a.Player.Stop)
	a.FyneApp.Lifecycle().SetOnExitedForeground(a.Player.Stop)
	a.main = a.FyneApp.NewWindow("Main - Animali")
	ContentChannal <- "Main"
	a.Main().ShowAndRun()
}

func (a *App) Main() fyne.Window {
	return a.main
}

func (a *App) SetContentWorker() {
	go func() {
		for val := range ContentChannal {
			fmt.Println(val)
			if screen, ok := a.Screen[val]; ok {
				a.main.SetContent(screen.Conteiner)
			}
		}
	}()
}
