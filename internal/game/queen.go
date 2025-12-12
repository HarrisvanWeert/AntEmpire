package game

import "time"

func StartQueen(g *GameState, ch GameChannels) {
	for {
		time.Sleep(2 * time.Second)

		if g.Food >= 5 {
			ch.StateChan <- StateEvent{FoodDelta: -5}
			ch.EggChan <- 1
			ch.LogChan <- "Queen laid Egg!"
		}
	}
}
