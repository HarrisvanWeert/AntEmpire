package assets

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type AntSprite struct {
	X, Y   float64
	VX, VY float64
}

func DrawAnts(screen *ebiten.Image, ants []AntSprite) {
	for _, ant := range ants {
		ebitenutil.DrawRect(screen, ant.X, ant.Y, 4, 4, color.White)
	}
}

// NewAntSprite returns a new random ant
func NewAntSprite() AntSprite {
	return AntSprite{
		X:  float64(100 + rand.Intn(400)),
		Y:  float64(100 + rand.Intn(300)),
		VX: float64(rand.Intn(3)+1) * 0.5,
		VY: float64(rand.Intn(3)+1) * 0.5,
	}
}
