package Extras

type EntityActions interface {
	attack()
	block()
	castSkill()
}
type Entity struct {
	HP          int
	DefenceStat int
	Energy      int
	Speed       int //higher speed stats mean that entity has turns earlier
	DodgeChance int
}
type Skill struct {
	Description string
}
type Weapon struct {
	weaponDamage int
	weaponSpeed  int
}

type Enemy struct {
	Entity
}
type Player struct {
	Entity
	Energy int
	weapon Weapon
}
