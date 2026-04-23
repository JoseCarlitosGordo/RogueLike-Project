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
	Active bool
}
type BossEncounter struct {
	Combat
}

func (g *Choice) Draw(gtx layout.Context, window *app.Window) {

}
