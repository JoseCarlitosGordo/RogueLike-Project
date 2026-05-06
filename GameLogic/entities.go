package GameLogic

var act1EnemyList = []EntityInCombat{CreateBrute(20, 3, 1, []SkillAction{CreateFlexSkill(2)}, 1, 5, 1.2, 1.3), CreateThief(5, 0, 5, []SkillAction{CreateMarkSkill(2)}, 1, 3, 10, 1.5)}

type EntityInCombat interface {
	CommenceTurn()
	Attack()
	Block(dmgReceived int)
	CastSkill(skill *Skill)
}

type Entity struct {
	HP          int
	DefenceStat int
	Speed       int //higher speed stats mean that entity has turns earlier
	Skills      []SkillAction
}

type Weapon struct {
	weaponDamage int
	weaponSpeed  int
}

type Enemy struct {
	Entity
	Damage               int
	Level                int
	RecentlyPlayedSkills []Skill
}

func (e *Enemy) CommenceTurn() {

}
func (e *Enemy) Attack() {

}

func (e *Enemy) Block(dmgReceived int) {
	dmgTaken := dmgReceived - e.DefenceStat
	e.HP -= dmgTaken
}

func (e *Enemy) CastSkill(s *Skill) {

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

func CreateThief(newHP int, newDefenceStat int, newSpeed int, skillList []SkillAction, level int, DamageValue int, CoinsStolenOnHit int, newSpeedMultiplier float32) *Thief {
	return &Thief{Enemy: Enemy{Entity: Entity{HP: newHP, DefenceStat: newDefenceStat, Skills: skillList}, Damage: DamageValue, Level: level}, NumberOfCoinsStolenOnHit: CoinsStolenOnHit, SpeedMultiplier: newSpeedMultiplier}
}

func CreateBrute(newHP int, newDefenceStat int, newSpeed int, skillList []SkillAction, level int, DamageValue int, newHealthMultiplier float32, newDamageMultiplier float32) *Brute {
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

func (p *Player) CommenceTurn() {

}
func (e *Player) Attack() {

}

func (e *Player) Block(dmgReceived int) {
	dmgTaken := dmgReceived - e.DefenceStat
	e.HP -= dmgTaken
}

func (e *Player) CastSkill(s *Skill) {

}
