package game

import (
	"time"
)

func StartWorker(ch GameChannels) {
	for {
		time.Sleep(1 * time.Second)
		ch.StateChan <- StateEvent{FoodDelta: 1}
		ch.LogChan <- "Worker gathered 1 food"
	}
}
