package main

import (
	"fmt"
	"harrisvw/internal/assets"
	"harrisvw/internal/game"
	"harrisvw/internal/ui"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	assets.LoadSprites()
	fmt.Println("Ant image size:", assets.Antimage.Bounds())

	ebiten.SetWindowTitle("Ant Colony")
	ebiten.SetWindowSize(960, 720)

	g := game.NewEbitenGame(ui.Draw)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Game Stopped")
}
