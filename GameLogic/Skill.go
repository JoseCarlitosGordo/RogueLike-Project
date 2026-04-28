package Extras

import (
	"fmt"

	"gioui.org/widget"
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

// // Test colors.
// var (
// 	background = color.NRGBA{R: 0xC0, G: 0xC0, B: 0xC0, A: 0xFF}
// 	red        = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
// 	black      = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
// 	green      = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: 0xFF}
// 	blue       = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}
// )

// // ColorBox creates a widget with the specified dimensions and color.
// func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
// 	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
// 	paint.ColorOp{Color: color}.Add(gtx.Ops)
// 	paint.PaintOp{}.Add(gtx.Ops)
// 	return layout.Dimensions{Size: size}
// }
// func (s *Skill) Draw(gtx layout.Context, MainTheme material.Theme) layout.Dimensions {
// 	return layout.Stack{}.Layout(gtx,
// 		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
// 			return layout.UniformInset(unit.Dp(30)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
// 				return ColorBox(gtx, gtx.Constraints.Max, red)
// 			})

// 		}),
// 		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
// 			return layout.UniformInset(unit.Dp(30)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
// 				return NewBody1(material.NewTheme(), s.Description, red, text.Middle).Layout(gtx)
// 			})
// 		}),
// 	)

// }

var markSkillInstance = CreateMarkSkill(1.3)

type HealSkill struct {
}
type ChargeSkill struct {
}
