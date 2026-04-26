package main

import (
	"Cool/Roguelike/Extras"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	//"gioui.org/widget/material"
)

var MainMenuState *Extras.MainMenu = Extras.CreateMainMenu()
var GameLoopState *Extras.GameState = &Extras.GameState{Active: false, IsStartingNewGame: false}

// mainMenu := Extras.MainMenu{active: true, theme: material.NewTheme(), title: GUIElement{text: "Hello!", style: theme}, PlayBtn: Button{{GUIElement{text: "Play", style: theme}, clickable: widget.Clickable}}}
func main() {
	//app logic sits in a goroutine
	go func() {
		MainMenuWindow := new(app.Window)
		err := run(MainMenuWindow)

		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	//Initialises application for certain OS's like MACOS and Android devices
	app.Main()
}

func run(window *app.Window) error {
	type Menu string
	const (
		Main_Menu    = "Main Menu"
		Options_Menu = "Options"
		Game_Loop    = "Game Loop"
		Pause_Menu   = "Pause Menu"
	)
	// options := Extras.CreateOptions()
	// gameLoop := Extras.BeginGame()
	// pauseMenu := Extras.PauseGame()
	//elementList := []Extras.GUIElement{}

	var ops op.Ops

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		//FrameEvent captures everything being displayed in gio
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			if MainMenuState.Active {
				MainMenuState.Draw(gtx, window, GameLoopState)
			}

			if GameLoopState.IsStartingNewGame {
				GameLoopState.BeginGame(gtx, MainMenuState.Theme)
			}

			if GameLoopState.Active {
				GameLoopState.Draw(gtx, MainMenuState.Theme)
			}

			e.Frame(gtx.Ops)
		}
	}
}
