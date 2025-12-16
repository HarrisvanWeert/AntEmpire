package game

import (
	"strconv"
	"time"
)

func StartNest(g *GameState, ch GameChannels) {
	workerID := g.Workers + 1

	for range ch.EggChan {
		select {
		case ch.LogChan <- "NEST: !!! Received Egg Signal !!!":
		default:
		}

		// Pause for 'hatching'
		time.Sleep(1 * time.Second)

		select {
		case ch.LogChan <- "NEST: Attempting to send StateEvent{WorkerDelta: 1}...":
		default:
		}

		select {
		case ch.StateChan <- StateEvent{WorkerDelta: 1}:
			// SUCCESSFUL SEND
			select {
			case ch.LogChan <- "NEST: SUCCESS! Worker StateEvent sent.":
			default:
			}
			select {
			case ch.LogChan <- "Egg hatched into worker :" + strconv.Itoa(workerID):
			default:
			}

			go StartWorker(ch)
			workerID++

		default:
			// FAILED SEND
			// This is the CRITICAL failure path.
			select {
			case ch.LogChan <- "NEST: FAILURE! StateChan is full. Worker NOT created.":
			default:
			}
		}
	}
}
