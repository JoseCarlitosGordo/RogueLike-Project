package Extras

import "gioui.org/widget/material"

var NavChoices []Choice = []Choice{
	{Active: false, ChoiceText: "A corridor splits in half up ahead. Do you go left or right?", choices: []Button{*CreateButtonStruct("Turn Left", material.NewTheme()), *CreateButtonStruct("Turn Right", material.NewTheme())}},
	{Active: false, ChoiceText: "A lamp lights up the room you are in. There are three closed doors...", choices: []Button{*CreateButtonStruct("Take the first door", material.NewTheme()), *CreateButtonStruct("Take the second door", material.NewTheme()), *CreateButtonStruct("Take the third door", material.NewTheme())}},
}
var EncounterChoices []Choice = []Choice{
	{Active: false, ChoiceText: "Ahead of you lies two chalices. The left chalice seems to grant great strength at the cost of some of your vitality while the right one gives you greater vitality at the cost of health..", choices: []Button{*CreateButtonStruct("Drink chalice on the left", material.NewTheme()), *CreateButtonStruct("Drink chalice on the right", material.NewTheme())}},
}

var VictoryChoice Choice = Choice{Active: false, ChoiceText: "You won the fight!", choices: []Button{*CreateButtonStruct("Press forward", material.NewTheme())}}
var BossChoice Choice = Choice{Active: false, ChoiceText: "You sense that a dastardly foe awaits ahead", choices: []Button{*CreateButtonStruct("No turning back now...", material.NewTheme())}}
