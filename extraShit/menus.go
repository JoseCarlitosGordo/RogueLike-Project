package extraShit

import (
	"image/color"

	"gioui.org/widget"
	"gioui.org/widget/material"
)

func NewButton(newButton *widget.Clickable, theme *material.Theme, text string) material.ButtonStyle {
	return material.Button(theme, newButton, text)
}

func NewText(theme *material.Theme, text string, color color.NRGBA) material.LabelStyle {
	return material.H1(theme, text)
}
