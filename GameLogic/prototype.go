package Extras

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	RoundsLeftUntilBoss := 5
	currentChoice := NavChoices[rand.IntN(len(NavChoices))]
	for RoundsLeftUntilBoss > 0 {
		fmt.Println(currentChoice)

		RoundsLeftUntilBoss -= 1

	}
}

func (combatEncounter *Combat) CommenceTurn() {

}
