package game

import (
	"harrisvw/internal/assets"
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
		StateChan: make(chan StateEvent, 50),
	}

	for i := 0; i < state.Workers; i++ {
		go StartWorker(ch)
	}

	go StartQueen(state, ch)
	go StartNest(state, ch)

	g := &EbitenGame{
		State:    state,
		Ch:       ch,
		lastTick: time.Now(),
		DrawUI:   drawUI, // Set it here
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

	for len(g.Ch.StateChan) > 0 {
		event := <-g.Ch.StateChan
		g.State.Food += event.FoodDelta
		g.State.Workers += event.WorkerDelta
	}

	//moving ants
	for i := range g.Ants {
		g.Ants[i].X += g.Ants[i].VX
		g.Ants[i].Y += g.Ants[i].VY

		if g.Ants[i].X < 0 || g.Ants[i].X > 640 {
			g.Ants[i].VX *= -1
		}
		if g.Ants[i].Y < 0 || g.Ants[i].Y > 480 {
			g.Ants[i].VY *= -1
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
