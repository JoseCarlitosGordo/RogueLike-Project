package main

import (
	"Cool/Roguelike/GameLogic"
)

func main() {
	MainGameState := GameLogic.GameState{RoundsLeftUntilBoss: 5, CurrentAct: 1}
	for MainGameState.CurrentAct <= 3 {
		MainGameState.RunSingleAct()
		MainGameState.CurrentAct += 1

	}

}
