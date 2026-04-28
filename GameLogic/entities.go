package GameLogic

type EntityActions interface {
	attack()
	block()
	castSkill()
}
type Entity struct {
	HP          int
	DefenceStat int
	Speed       int //higher speed stats mean that entity has turns earlier
	Skills      []Skill
}

type Weapon struct {
	weaponDamage int
	weaponSpeed  int
}

type Enemy struct {
	Entity
	Damage int
	Level  int
}

type Boss struct {
	Enemy
	HealthMultiplier int
}

type Thief struct {
	Enemy
	NumberOfCoinsStolenOnHit int

	SpeedMultiplier float32
}
type Brute struct {
	Enemy
	HealthMultiplier float32
	DamageMultiplier float32
}

func CreateThief(newHP int, newDefenceStat int, newSpeed int, skillList []Skill, level int, DamageValue int, CoinsStolenOnHit int, newSpeedMultiplier float32) *Thief {
	return &Thief{Enemy: Enemy{Entity: Entity{HP: newHP, DefenceStat: newDefenceStat, Skills: skillList}, Damage: DamageValue, Level: level}, NumberOfCoinsStolenOnHit: CoinsStolenOnHit, SpeedMultiplier: newSpeedMultiplier}
}

func CreateBrute(newHP int, newDefenceStat int, newSpeed int, skillList []Skill, level int, DamageValue int, newHealthMultiplier float32, newDamageMultiplier float32) *Brute {
	return &Brute{Enemy: Enemy{Entity: Entity{HP: newHP, DefenceStat: newDefenceStat, Skills: skillList}, Damage: DamageValue, Level: level}, DamageMultiplier: newDamageMultiplier, HealthMultiplier: newHealthMultiplier}
}

type Player struct {
	Entity
	MaxEnergy  int
	Energy     int
	MoneyCount int
	Weapon     Weapon
	LuckStat   int
}
