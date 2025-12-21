package assets

import (
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var Antimage *ebiten.Image
var BackGroundImage *ebiten.Image

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

	bgFile, err := os.Open("../../internal/assets/images/terrain.png")
	if err != nil {
		panic(err)
	}
	defer bgFile.Close()

	bgImage, err := png.Decode(bgFile)
	if err != nil {
		panic(err)
	}

	BackGroundImage = ebiten.NewImageFromImage(bgImage)
	Antimage = ebiten.NewImageFromImage(img)

}
