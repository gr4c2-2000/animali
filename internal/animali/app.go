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
	a.MusicScreenResources = []fyne.Resource{resourceMusic1Png, resourceMusic2Png, resourceMusic3Png, resourceMusic4Png, resourceMusic5Png, resourceMusic6Png, resourceMusic7Png, resourceMusic8Png, resourceMusic9Png}
	a.Player = InitPayer()
	a.Player.AddToPlaylist(resourceSongMp3, resourceSongMp3.StaticName)
	a.Player.AddToPlaylist(resourceShortMp3, resourceShortMp3.StaticName)
	a.SetContentWorker()
	return &a
}

func (a *App) Run() {
	a.FyneApp = app.New()
	a.FyneApp.Preferences()
	a.FyneApp.Settings().SetTheme(&myTheme{})
	mv := BuildMainView()
	a.Screen = make(map[string]Screen)
	MainScr := Screen{title: "Main", Conteiner: mv.container}
	a.Screen["Main"] = MainScr
	Mv := Build(a.Player)
	Music := Screen{title: "Music", Conteiner: Mv.Container()}
	a.Screen["Music"] = Music
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
				fmt.Println("Set")
				a.main.SetContent(screen.Conteiner)
			}
		}
	}()
}
