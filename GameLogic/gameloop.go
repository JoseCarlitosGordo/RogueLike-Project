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
	Active    bool
	EnemyList []Enemy
}
type BossFight struct {
	Combat
}
type RandomEncounter struct {
	Active        bool
	EncounterText string
	Choices       Choice
}

func (c *Combat) RunEvent() {

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
			chosenChoice.consequence = &Combat{EnemyList: []Enemy{}}
		} else {

		}
		chosenChoice.consequence.RunEvent()

		g.RoundsLeftUntilBoss -= 1

	}

}
func (g *GameState) CommenceBossFight() {

}
