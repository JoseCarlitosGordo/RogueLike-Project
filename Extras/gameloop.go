package Extras

import (
	"image"
	"image/color"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget/material"
)

type GameComponent interface {
	CreateList()
}
type Choice struct {
	Active     bool
	ChoiceText string
	choices    []Button
}
type Combat struct {
	Active    bool
	EnemyList []Enemy
}
type BossFight struct {
	Combat
}
type RandomEncounter struct {
	Active        bool
	EncounterText string
	Choices       []Button
}

func (g *Choice) Draw(gtx layout.Context, window *app.Window) {

}

func (g *Combat) Draw(gtx layout.Context, window *app.Window) {

}

type GameState struct {
	Active bool

	IsStartingNewGame bool
	// Choice
	// randomEncounters []RandomEncounter

}

var buttonData = CreateButtonStruct("Click here to proceed", material.NewTheme())

func (g *GameState) GiveMonologue(gtx layout.Context, MainTheme *material.Theme) layout.Dimensions {
	title := NewH4(MainTheme, "A new Adventure begins", color.NRGBA{R: 127, G: 0, B: 0, A: 255}, text.Middle)
	paint.Fill(gtx.Ops, color.NRGBA{R: 0, G: 0, B: 0, A: 255})
	proceedBtn := NewButton(&buttonData.clickable, buttonData.style, buttonData.text)
	if proceedBtn.Button.Clicked(gtx) {
		g.IsStartingNewGame = false
		g.Active = true
	}

	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Middle,
		Spacing:   layout.SpaceEvenly,
	}.Layout(gtx, layout.Rigid(title.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// Set both Min and Max to force an exact width
			gtx.Constraints.Min.X = gtx.Dp(200)
			gtx.Constraints.Max.X = gtx.Dp(200)
			return proceedBtn.Layout(gtx)
		}),
	)

}
func (g *GameState) BeginGame(gtx layout.Context, MainTheme *material.Theme) {
	g.GiveMonologue(gtx, MainTheme)

}

// Test colors.
var (
	background = color.NRGBA{R: 0xC0, G: 0xC0, B: 0xC0, A: 0xFF}
	red        = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
	green      = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: 0xFF}
	blue       = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}
)

// ColorBox creates a widget with the specified dimensions and color.
func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}
func (g *GameState) Draw(gtx layout.Context, MainTheme *material.Theme) layout.Dimensions {
	title := NewH4(material.NewTheme(), "A new Adventure begins 2", color.NRGBA{R: 0, G: 0, B: 0, A: 255}, text.Middle)

	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Middle,
		Spacing:   layout.SpaceEvenly,
	}.Layout(gtx, layout.Rigid(title.Layout))

}
