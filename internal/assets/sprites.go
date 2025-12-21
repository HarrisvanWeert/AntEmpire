package assets

import (
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var Antimage *ebiten.Image

func LoadSprites() {
	file, err := os.Open("../../internal/assets/images/ant.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	Antimage = ebiten.NewImageFromImage(img)

}
