package ui

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"

	"harrisvw/internal/game"
)

func Draw(screen *ebiten.Image, state *game.GameState) {
	text.Draw(screen, "ANT COLONY SIM", basicfont.Face7x13, 20, 30, color.White)
	text.Draw(screen, fmt.Sprintf("Tick: %d", state.Tick), basicfont.Face7x13, 20, 50, color.White)
	text.Draw(screen, fmt.Sprintf("Ants: %d", state.Workers), basicfont.Face7x13, 20, 70, color.White)
	text.Draw(screen, fmt.Sprintf("Food: %d", state.Food), basicfont.Face7x13, 20, 90, color.White)
	text.Draw(screen, fmt.Sprintf("Eggs: %d", state.Eggs), basicfont.Face7x13, 20, 110, color.White)
}
