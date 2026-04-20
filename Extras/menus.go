package Extras

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func DisplayMainMenu(appContext *layout.Context, theme *material.Theme, elementsToDisplay []GUIElement) {

}

func NewButton(newButton *widget.Clickable, theme *material.Theme, text string) material.ButtonStyle {
	return material.Button(theme, newButton, text)
}

func NewText(theme *material.Theme, text string, color color.NRGBA) material.LabelStyle {
	newText := material.H1(theme, text)
	newText.Color = color
	return newText
}
