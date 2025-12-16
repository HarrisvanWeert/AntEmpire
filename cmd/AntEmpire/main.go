package main

import (
	"log"

	"harrisvw/internal/game"
	"harrisvw/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Ant Colony")
	ebiten.SetWindowSize(960, 720)

	g := game.NewEbitenGame()
	g.DrawUI = ui.Draw

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
