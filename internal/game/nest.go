package game

import (
	"strconv"
	"time"
)

func StartNest(g *GameState, ch GameChannels) {
	workerID := g.Workers + 1

	for range ch.EggChan {
		time.Sleep(1 * time.Second)

		ch.StateChan <- StateEvent{WorkerDelta: 1}
		ch.LogChan <- "Egg hatched into worker :" + strconv.Itoa(workerID)

		go StartWorker(ch)
		workerID++
	}
}
