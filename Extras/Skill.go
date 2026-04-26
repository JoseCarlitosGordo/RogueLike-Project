package Extras

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type SkillAction interface {
	PerformSkill()
	Upgrade()
}

type Skill struct {
	clickable   widget.Clickable
	Description string
	ManaCost    int
}
type MarkSkill struct {
	DamageMultiplier float32
	Skill
}

func CreateMarkSkill(newMultiplier float32) *MarkSkill {
	newDescription := fmt.Sprintf("Target an enemy. When Targeted, they take %v percent more damage from any attack", (newMultiplier-1)*100)
	return &MarkSkill{DamageMultiplier: newMultiplier, Skill: Skill{clickable: widget.Clickable{}, Description: newDescription, ManaCost: 2}}

}

func (s *Skill) Draw(gtx layout.Context, MainTheme material.Theme) layout.Dimensions {
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(30)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return ColorBox(gtx, gtx.Constraints.Max, red)
			})

		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(30)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return NewBody1(material.NewTheme(), s.Description, red, text.Middle).Layout(gtx)
			})
		}),
	)

}

var markSkillInstance = CreateMarkSkill(1.3)

type HealSkill struct {
}
type ChargeSkill struct {
}
