package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	// the image and position variables for our player
	PlayerImg *ebiten.Image
	X, Y      float64
}

func (g *Game) Update() error {

	// move the player based on keyboar input (left, right, up down)
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Y += 2
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// fill the screen with a nice sky color
	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}
	// set the translation of our drawImageOptions to the player's position
	opts.GeoM.Translate(g.X, g.Y)

	// draw the player
	screen.DrawImage(
		// grab a subimage of the spritesheet
		g.PlayerImg.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image),
		&opts,
	)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// load the image from file
	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/ninja.png")
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	game := Game{
		PlayerImg: playerImg,
		X:         0.0,
		Y:         0.0,
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
