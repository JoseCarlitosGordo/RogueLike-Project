package Extras

import (
	"image"
	"image/color"
	"math/rand/v2"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func (buttonChoice *Button) CreateAndRunConsequence() {
	//TODO: If num is greater than 3, ChoiceConsequence is a combat encounter. Otherwise, it is a random encounter

}
func (choice *Choice) Draw(gtx layout.Context, th *material.Theme) layout.Dimensions {
	title := NewH4(th, "Make a choice....", color.NRGBA{R: 0, G: 0, B: 0, A: 255}, text.Middle)
	physicalChoiceButtons := layout.List{Axis: layout.Horizontal}
	buttonLayout := physicalChoiceButtons.Layout(gtx, len(choice.choices),
		func(gtx layout.Context, i int) layout.Dimensions {
			btn := NewButton(&choice.choices[i].clickable, choice.choices[i].style, choice.choices[i].text)
			if btn.Button.Clicked(gtx) {
				choice.choices[i].CreateAndRunConsequence()

			}
			return btn.Layout(gtx)
		})
	return layout.Flex{Alignment: layout.Middle}.Layout(gtx, layout.Rigid(title.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return buttonLayout
		}),
	)

}

func (combat *Combat) Draw(gtx layout.Context) {

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
	var layout_to_return layout.Dimensions = layout.Dimensions{Size: image.Point{}}
	if g.IsMakingAChoice {
		layout_to_return = g.DisplayNavChoice(gtx, MainTheme)
	}

	return layout_to_return

}

func (g *GameState) DisplayNavChoice(gtx layout.Context, MainTheme *material.Theme) layout.Dimensions {
	newChoice := NavChoices[rand.IntN(len(NavChoices))]
	return newChoice.Draw(gtx, MainTheme)

}
