package extraShit

import (
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Buttons struct {
	button widget.Clickable
	theme  *material.Theme
}

func NewButton (newButton widget.Clickable, theme *material.Theme) *widget.Clickable
{
	
}