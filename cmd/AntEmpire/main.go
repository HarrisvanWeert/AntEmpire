package main

import (
	"harrisvw/internal/game"
	"harrisvw/internal/ui"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Ant Colony")
	ebiten.SetWindowSize(960, 720)

	g := game.NewEbitenGame(ui.Draw)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
