package game

import "time"

func StartQueen(ch GameChannels) {
	for {
		time.Sleep(5 * time.Second)
		reply := make(chan int)
		ch.FoodQuery <- reply
		currentFood := <-reply

		if currentFood >= 2 {
			ch.StateChan <- StateEvent{FoodDelta: -2}
			ch.EggChan <- 1
			select {
			case ch.LogChan <- "Queen laid Egg!":
			default:
			}
		}
	}
}
