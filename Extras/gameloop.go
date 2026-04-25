package Extras

import (
	"gioui.org/app"
	"gioui.org/layout"
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
	Active    bool
	IsNewGame bool
	// Choice
	// randomEncounters []RandomEncounter

}

func GiveMonologue() {

}
func (g *GameState) BeginGame() {
	GiveMonologue()

}

func RunGame() {

}
