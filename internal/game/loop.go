package game

import (
	"fmt"
	"harrisvw/internal/ui"
	"time"
)

func GameLoop() {

	game := &GameState{
		Food:    0,
		Workers: 3,
		Eggs:    0,
	}

	ch := GameChannels{
		EggChan:   make(chan int),
		LogChan:   make(chan string, 20),
		StateChan: make(chan StateEvent, 50),
	}

	for i := 0; i < game.Workers; i++ {
		go StartWorker(ch)
	}

	go StartQueen(game, ch)
	go StartNest(game, ch)

	ticker := time.NewTicker(250 * time.Millisecond)
	defer ticker.Stop()

	rollingLogs := []string{}
	const maxLogs = 8

	for range ticker.C {
		game.Tick++

		for len(ch.StateChan) > 0 {
			event := <-ch.StateChan
			game.Food += event.FoodDelta
			game.Workers += event.WorkerDelta
		}

		for len(ch.LogChan) > 0 {
			rollingLogs = append(rollingLogs, <-ch.LogChan)

			if len(rollingLogs) > maxLogs {
				rollingLogs = rollingLogs[1:]
			}
		}

		ui.Clear()

		fmt.Println("Events:")
		for _, log := range rollingLogs {
			fmt.Println(" -", log)
		}
		fmt.Println()

		ui.DisplayStats(game.Workers, game.Food, game.Tick)
	}
}
