package Extras

import (
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// logic for holding elements inside other structs
type GUIElement struct {
	//elementType T
	text  string
	style *material.Theme
}

type Button struct {
	GUIElement
	clickable widget.Clickable
}

func CreateGUIElementStruct(newText string, newStyle *material.Theme) *GUIElement {
	return &GUIElement{text: newText, style: newStyle}
}
func CreateButtonStruct(newText string, newStyle *material.Theme) *Button {
	return &Button{GUIElement: *CreateGUIElementStruct(newText, newStyle), clickable: widget.Clickable{}}
}
