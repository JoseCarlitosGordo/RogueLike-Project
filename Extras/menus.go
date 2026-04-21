package Extras

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func DisplayMainMenu(appContext *layout.Context, theme *material.Theme, elementsToDisplay []GUIElement) {

}

func NewButton(newButton *widget.Clickable, theme *material.Theme, text string) material.ButtonStyle {
	return material.Button(theme, newButton, text)
}

func NewText(theme *material.Theme, text string, color color.NRGBA) material.LabelStyle {
	newText := material.H1(theme, text)
	newText.Color = color
	return newText
}

type Display interface {
	draw()
}

type MainMenu struct {
	active     bool
	theme      *material.Theme
	title      *GUIElement
	PlayBtn    *Button //{GUIElement{text: "Play", style: theme}, clickable: widget.Clickable}
	OptionsBtn *Button
	QuitBtn    *Button
}

func CreateMainMenu() *MainMenu {
	newTheme := material.NewTheme()
	return &MainMenu{active: true, theme: newTheme, title: CreateGUIElement("Working title for Roguelike", newTheme), PlayBtn: CreateButton("Play", newTheme), OptionsBtn: CreateButton("Options", newTheme), QuitBtn: CreateButton("Quit", newTheme)}

}
func (m *MainMenu) draw() {
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

}

type Options struct {
	active bool
}
type GameLoop struct {
	active bool
}
