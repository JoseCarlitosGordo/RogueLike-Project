package Extras

import "fmt"

type SkillActions interface {
	PerformSkill()
	Upgrade()
}

type Skill struct {
	Description string
}
type MarkSkill struct {
	multiplier float32
	Skill
}

func CreateMarkSkill(newMultiplier float32) *MarkSkill {
	newDescription := fmt.Sprintf("Target an enemy. When Targeted, they take %v percent more damage from any attack", (newMultiplier-1)*100)
	return &MarkSkill{multiplier: newMultiplier, Skill: Skill{Description: newDescription}}

}

var markSkillInstance = CreateMarkSkill(1.3)

type HealSkill struct {
}
type ChargeSkill struct {
}
