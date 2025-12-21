package assets

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type AntSprite struct {
	X, Y   float64
	VX, VY float64
}

func DrawAnts(screen *ebiten.Image, ants []AntSprite) {
	for _, ant := range ants {
		op := &ebiten.DrawImageOptions{}

		scale := 0.02
		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(ant.X, ant.Y)

		screen.DrawImage(Antimage, op)
	}
}

func NewAntSprite() AntSprite {
	randomVelocity := func() float64 {
		v := (rand.Float64() * 4) - 2
		if v > -0.5 && v < 0.5 {
			if v >= 0 {
				v = 0.5 + rand.Float64()*1.5
			} else {
				v = -0.5 - rand.Float64()*1.5
			}
		}
		return v
	}

	return AntSprite{
		X:  float64(rand.Intn(540) + 50),
		Y:  float64(rand.Intn(380) + 50),
		VX: randomVelocity(),
		VY: randomVelocity(),
	}
}
