package Extras

import (
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type GUIElement struct {
	//elementType T
	text  string
	style *material.Theme
}

type Button struct {
	GUIElement
	clickable widget.Clickable
}

func CreateGUIElement(newText string, newStyle *material.Theme) *GUIElement {
	return &GUIElement{text: newText, style: newStyle}
}
func CreateButton(newText string, newStyle *material.Theme) *Button {
	return &Button{GUIElement: *CreateGUIElement(newText, newStyle), clickable: widget.Clickable{}}
}
