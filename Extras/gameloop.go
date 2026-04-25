package Extras

import (
	"image/color"

	"gioui.org/app"
	"gioui.org/layout"
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

type GameStart struct {
}
type GameState struct {
	Active bool

	IsStartingNewGame bool
	// Choice
	// randomEncounters []RandomEncounter

}

func (g *GameState) GiveMonologue(gtx layout.Context) layout.Dimensions {
	title := NewH4(material.NewTheme(), "A new Adventure begins", color.NRGBA{R: 0, G: 0, B: 0, A: 255}, text.Middle)
	buttonData := CreateButtonStruct("Click here to proceed", material.NewTheme())
	proceedBtn := NewButton(&buttonData.clickable, buttonData.style, buttonData.text)
	if proceedBtn.Button.Clicked(gtx) {
		g.IsStartingNewGame = false
		g.Active = true
	}
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
			return proceedBtn.Layout(gtx)
		}),
	)

}
func (g *GameState) BeginGame(gtx layout.Context) {
	g.GiveMonologue(gtx)

}

func (g *GameState) Draw(gtx layout.Context) layout.Dimensions {
	title := NewH4(material.NewTheme(), "A new Adventure begins 2", color.NRGBA{R: 255, G: 255, B: 255, A: 255}, text.Middle)

	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Middle,
		Spacing:   layout.SpaceEvenly,
	}.Layout(gtx, layout.Rigid(title.Layout))

}
