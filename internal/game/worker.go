package game

import (
	"time"
)

func StartWorker(ch GameChannels) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case ch.StateChan <- StateEvent{FoodDelta: 1}:
			select {
			case ch.LogChan <- "Worker gathered 1 food":
			default:
			}
		default:
		}
	}
}
