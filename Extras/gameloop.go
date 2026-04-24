package Extras

import (
	"gioui.org/app"
	"gioui.org/layout"
)

type Choice struct {
	Active   bool
	LeftBtn  Button
	RightBtn Button
}
type Combat struct {
	Active    bool
	EnemyList []Enemy
}
type BossEncounter struct {
	Combat
}
type RandomEncounter struct {
	EncounterText string
	Choices       []Button
}

func (g *Choice) Draw(gtx layout.Context, window *app.Window) {

}

func (g *Combat) Draw(gtx layout.Context, window *app.Window) {

}
