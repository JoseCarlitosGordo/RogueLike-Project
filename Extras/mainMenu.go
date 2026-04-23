package Extras

import (
	"image/color"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func NewButton(newButton *widget.Clickable, theme *material.Theme, text string) material.ButtonStyle {
	return material.Button(theme, newButton, text)
}

func NewH4(theme *material.Theme, text string, color color.NRGBA, alignment text.Alignment) material.LabelStyle {
	newText := material.H4(theme, text)
	newText.Color = color
	newText.Alignment = alignment
	return newText
}

type Display interface {
	draw()
}

type MainMenu struct {
	Active     bool
	title      *GUIElement
	PlayBtn    *Button //{GUIElement{text: "Play", style: theme}, clickable: widget.Clickable}
	OptionsBtn *Button
	QuitBtn    *Button
}

func CreateMainMenu() *MainMenu {
	newTheme := material.NewTheme()
	return &MainMenu{Active: true, title: CreateGUIElement("Working title for Roguelike", newTheme), PlayBtn: CreateButton("Play", newTheme), OptionsBtn: CreateButton("Options", newTheme), QuitBtn: CreateButton("Quit", newTheme)}

}
func (m *MainMenu) Draw(gtx layout.Context, window *app.Window) layout.Dimensions {
	title := NewH4(m.title.style, m.title.text, color.NRGBA{R: 127, G: 0, B: 0, A: 255}, text.Middle)
	bgColor := color.NRGBA{R: 0, G: 0, B: 0, A: 255} // Light red
	playBtn := NewButton(&m.PlayBtn.clickable, m.PlayBtn.style, m.PlayBtn.text)
	// Fill the background
	paint.Fill(gtx.Ops, bgColor)
	optionsBtn := NewButton(&m.OptionsBtn.clickable, m.OptionsBtn.style, m.OptionsBtn.text)
	quitBtn := NewButton(&m.QuitBtn.clickable, m.QuitBtn.style, m.QuitBtn.text)
	if playBtn.Button.Clicked(gtx) {
		//

	}

	if optionsBtn.Button.Clicked(gtx) {

	}

	if quitBtn.Button.Clicked(gtx) {
		window.Perform(system.ActionClose)
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
	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Middle,
		Spacing:   layout.SpaceEvenly,
	}.Layout(gtx,
		layout.Rigid(title.Layout),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// Set both Min and Max to force an exact width
			gtx.Constraints.Min.X = gtx.Dp(200)
			gtx.Constraints.Max.X = gtx.Dp(200)
			return playBtn.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// Set both Min and Max to force an exact width
			gtx.Constraints.Min.X = gtx.Dp(200)
			gtx.Constraints.Max.X = gtx.Dp(200)
			return optionsBtn.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// Set both Min and Max to force an exact width
			gtx.Constraints.Min.X = gtx.Dp(200)
			gtx.Constraints.Max.X = gtx.Dp(200)
			return quitBtn.Layout(gtx)
		}),
	)

}
