package Extras

import (
	"image/color"
	"math/rand/v2"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget/material"
)

type GameComponent interface {
	CreateList()
}
type Choice struct {
	Active     bool
	Repeatable bool
	ChoiceText string
	choices    []Button
}

type ChoiceConsequence struct {
	ChoiceText string
}
type Combat struct {
	ChoiceConsequence
	Active    bool
	EnemyList []Enemy
}
type BossFight struct {
	Combat
}
type RandomEncounter struct {
	ChoiceConsequence
	Active        bool
	EncounterText string
	Choices       Choice
}

func (g *Choice) Draw(gtx layout.Context, th *material.Theme) layout.Dimensions {
	for i := range g.choices {
		btn := NewButton(&g.choices[i].clickable, g.choices[i].style, g.choices[i].text)
		btn.Layout(gtx)

	}
	return layout.Flex{}.Layout(gtx)

}

func (g *Combat) Draw(gtx layout.Context) {

}

type GameState struct {
	Active bool

	IsStartingNewGame   bool
	IsMakingAChoice     bool
	RoundCount          int
	CurrentChoice       []Choice
	PotentialEncounters []RandomEncounter
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
func (g *GameState) Draw(gtx layout.Context, MainTheme *material.Theme) layout.Dimensions {
	title := NewH4(MainTheme, "Make a choice....", color.NRGBA{R: 0, G: 0, B: 0, A: 255}, text.Middle)
	if g.IsMakingAChoice {
		g.DisplayNavChoice(gtx, MainTheme)

	}
	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Middle,
		Spacing:   layout.SpaceEvenly,
	}.Layout(gtx, layout.Rigid(title.Layout))

}

func (g *GameState) DisplayNavChoice(gtx layout.Context, MainTheme *material.Theme) layout.Dimensions {
	newChoice := NavChoices[rand.IntN(len(NavChoices))]
	return newChoice.Draw(gtx, MainTheme)

}
