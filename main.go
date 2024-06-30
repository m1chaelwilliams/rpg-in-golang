// the package our code is scoped to. The `main` package is where the program's entry point lives
package main

// All imports we are using in this file
import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Our struct that impliments the `ebiten.Game` interface
// No `impliments x` needed in Golang, the compiler does that for you!
type Game struct{}

// where all of our update functionality will live
func (g *Game) Update() error {
	return nil
}

// where all of our draw functionality will live
func (g *Game) Draw(screen *ebiten.Image) {

	// fills the screen with a light blue color
	screen.Fill(color.RGBA{120, 180, 255, 255})

	// draws "Hello, World!" on the screen
	ebitenutil.DebugPrint(screen, "Hello, World!")

}

// determines the dimensions of our rendering canvas (not the window size!)
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// the entry point of our application
func main() {
	// setting various parameters of the app
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// running the game and handling any errors that are routed up to us
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
