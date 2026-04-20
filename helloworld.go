package main

import (
	"Cool/Roguelike/Extras"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		MainMenuWindow := new(app.Window)
		err := run(MainMenuWindow)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	type Menu string
	const (
		Main_Menu    = "Main Menu"
		Options_Menu = "Options"
		Game_Loop    = "Game Loop"
	)
	currentMenu := Main_Menu
	theme := material.NewTheme()
	elementList := []Extras.GUIElement{}

	var ops op.Ops

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		//FrameEvent captures everything being displayed in gio
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)
			switch currentMenu {
			case Main_Menu:
				Extras.DisplayMainMenu(&gtx, theme, elementList)
			case Options_Menu:

			case Game_Loop:

			}

			// // Define an large label with an appropriate text:
			// title := material.H1(theme, "Hello, Gio")
			// // Change the color of the label.
			// maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			// title.Color = maroon
			// // Change the position of the label.
			// title.Alignment = text.Middle

			// // Draw the label to the graphics context.
			// title.Layout(gtx)

			// PlayBtn := extraShit.NewButton(&mainMenuButton, theme, "Play")
			// PlayBtn.Layout(gtx)
			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
