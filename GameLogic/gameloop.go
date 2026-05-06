package GameLogic

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

type GameState struct {
	RoundsLeftUntilBoss int
	CurrentAct          int
	CombatChance        int
}
type Event interface {
	RunEvent()
}
type Combat struct {
	Active      bool
	EnemyList   []EntityInCombat
	player      *Player
	CurrentTurn EntityInCombat
}
type BossFight struct {
	Combat
}
type RandomEncounter struct {
	Active        bool
	EncounterText string
	Choices       Choice
}

func (rm *RandomEncounter) RunEvent() {

}
func (c *Combat) RunEvent() {
	enemy_number := rand.IntN(3) + 2
	for i := range enemy_number {
		c.EnemyList = append(c.EnemyList, act1EnemyList[rand.IntN(len(act1EnemyList))])
		i += 1
	}

	for c.player.HP > 0 && len(c.EnemyList) > 0 {

		c.ProcessTurn(c.player)

		for _, enemy := range c.EnemyList {
			c.ProcessTurn(enemy)
		}
	}
}
func CreateEnemies(enemyList []EntityInCombat)

func (c *Combat) ProcessTurn(entity EntityInCombat) {
	entity.CommenceTurn()

}

func (g *GameState) RunSingleAct() {

	g.RoundsLeftUntilBoss = 5
	currentChoice := NavChoices[rand.IntN(len(NavChoices))]
	input := ""
	for g.RoundsLeftUntilBoss > 0 {
		fmt.Println(currentChoice)

		for _, option := range currentChoice.choices {
			fmt.Println(option.text)
		}
		fmt.Scanln(&input)
		optionChosen, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("smth went horribly wrong near line 28 of gameloop.go: %v\n", err.Error())
			return
		}
		chosenChoice := currentChoice.choices[optionChosen-1]
		if rand.IntN(101) < g.CombatChance {
			chosenChoice.consequence = &Combat{}
		} else {
			chosenChoice.consequence = &RandomEncounter{}
		}
		chosenChoice.consequence.RunEvent()

		g.RoundsLeftUntilBoss -= 1

	}

}
func (b *BossFight) RunEvent() {

}
