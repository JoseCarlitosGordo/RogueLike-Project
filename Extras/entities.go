package Extras

type EntityActions interface {
	attack()
	block()
	castSkill()
}
type Entity struct {
	HP          int
	DefenceStat int
	Speed       int //higher speed stats mean that entity has turns earlier
	DodgeChance int
	Skills      []Skill
	Exp         int
}

type Weapon struct {
	weaponDamage int
	weaponSpeed  int
}

type Enemy struct {
	Entity
	Damage int
}
type Player struct {
	Entity
	Energy int
	Weapon Weapon
}
