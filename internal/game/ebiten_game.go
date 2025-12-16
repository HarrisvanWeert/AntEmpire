package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// EbitenGame wraps your game state and channels
type EbitenGame struct {
	State *GameState
	Ch    GameChannels

	lastTick time.Time

	// DrawUI is a callback for rendering, set from main
	DrawUI func(screen *ebiten.Image, state *GameState)
	logs   []string
}

// NewEbitenGame initializes the game and starts goroutines
func NewEbitenGame() *EbitenGame {
	state := &GameState{
		Food:    0,
		Workers: 3,
		Eggs:    0,
	}

	ch := GameChannels{
		EggChan:   make(chan int),
		LogChan:   make(chan string, 20),
		StateChan: make(chan StateEvent, 50),
	}

	for i := 0; i < state.Workers; i++ {
		go StartWorker(ch)
	}

	go StartQueen(state, ch)
	go StartNest(state, ch)

	return &EbitenGame{
		State:    state,
		Ch:       ch,
		lastTick: time.Now(),
	}
}

// Update advances the game logic every 250ms
func (g *EbitenGame) Update() error {
	if time.Since(g.lastTick) < 250*time.Millisecond {
		return nil
	}
	g.lastTick = time.Now()

	g.State.Tick++

	for len(g.Ch.StateChan) > 0 {
		event := <-g.Ch.StateChan
		g.State.Food += event.FoodDelta
		g.State.Workers += event.WorkerDelta
	}

	return nil
}

// Draw calls the UI callback
func (g *EbitenGame) Draw(screen *ebiten.Image) {
	if g.DrawUI != nil {
		g.DrawUI(screen, g.State)
	}
}

// Layout sets the window size
func (g *EbitenGame) Layout(w, h int) (int, int) {
	return 640, 480
}
