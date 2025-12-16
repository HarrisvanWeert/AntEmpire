package game

import (
	"harrisvw/internal/assets"
	"math/rand/v2"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type EbitenGame struct {
	State    *GameState
	Ch       GameChannels
	lastTick time.Time
	DrawUI   func(screen *ebiten.Image, state *GameState, ants []assets.AntSprite)
	Ants     []assets.AntSprite
	logs     []string
}

func NewEbitenGame(drawUI func(screen *ebiten.Image, state *GameState, ants []assets.AntSprite)) *EbitenGame {
	state := &GameState{
		Food:    0,
		Workers: 3,
		Eggs:    0,
	}

	ch := GameChannels{
		EggChan:   make(chan int),
		LogChan:   make(chan string, 20),
		StateChan: make(chan StateEvent, 5000),
		FoodQuery: make(chan chan int),
	}

	for i := 0; i < state.Workers; i++ {
		go StartWorker(ch)
	}

	go StartQueen(ch)
	go StartNest(state, ch)

	g := &EbitenGame{
		State:    state,
		Ch:       ch,
		lastTick: time.Now(),
		DrawUI:   drawUI,
	}

	for i := 0; i < state.Workers; i++ {
		g.Ants = append(g.Ants, assets.NewAntSprite())
	}

	return g
}

func (g *EbitenGame) Update() error {
	if time.Since(g.lastTick) < 250*time.Millisecond {
		return nil
	}
	g.lastTick = time.Now()

	g.State.Tick++

	for {
		select {
		case msg := <-g.Ch.LogChan:
			g.logs = append(g.logs, msg)
			if len(g.logs) > 100 {
				g.logs = g.logs[len(g.logs)-100:]
			}
		default:
			goto doneDrainingLogs
		}
	}
doneDrainingLogs:

	select {
	case reply := <-g.Ch.FoodQuery:
		reply <- g.State.Food
	default:
	}
	for {
		select {
		case event := <-g.Ch.StateChan:
			g.State.Food += event.FoodDelta
			g.State.Workers += event.WorkerDelta

			if event.WorkerDelta > 0 {
				g.Ants = append(g.Ants, assets.NewAntSprite())
			}
		default:
			goto doneProcessing
		}
	}
doneProcessing:

	for i := range g.Ants {
		g.Ants[i].X += g.Ants[i].VX
		g.Ants[i].Y += g.Ants[i].VY

		if g.Ants[i].X < 0 || g.Ants[i].X > 640 {
			g.Ants[i].VX *= -1
		}
		if g.Ants[i].Y < 0 || g.Ants[i].Y > 480 {
			g.Ants[i].VY *= -1
		}
		if rand.Float64() < 0.05 {
			g.Ants[i].VX += (rand.Float64() - 0.5) * 0.5
			g.Ants[i].VY += (rand.Float64() - 0.5) * 0.5

			if g.Ants[i].VX > 3 {
				g.Ants[i].VX = 3
			} else if g.Ants[i].VX < -3 {
				g.Ants[i].VX = -3
			}
			if g.Ants[i].VY > 3 {
				g.Ants[i].VY = 3
			} else if g.Ants[i].VY < -3 {
				g.Ants[i].VY = -3
			}
		}
	}

	return nil
}

func (g *EbitenGame) Draw(screen *ebiten.Image) {
	if g.DrawUI != nil {
		g.DrawUI(screen, g.State, g.Ants)
	}
}

func (g *EbitenGame) Layout(w, h int) (int, int) {
	return 640, 480
}
