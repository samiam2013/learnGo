package main

import (
	"image/color"
	"log"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game interface.
type Game struct {
	ball          ball
	groundPos     int
	ground        *ebiten.Image
	groundOptions *ebiten.DrawImageOptions
}

type ball struct {
	radius       int
	vSpeed       float32
	hSpeed       float32
	vPos         int
	hPos         int
	image        *ebiten.Image
	imageOptions *ebiten.DrawImageOptions
}

// applyGravity updates the position and speed of the ball for a single tick
func (b *ball) applyGravity() {
	// make the change according to the current speed

	// computer the change in speed for 1/60 of a second

}

func (b *ball) makeImage() *ebiten.Image {
	//ball := ebiten.NewImage(b.radius*2, b.radius*2)
	dc := gg.NewContext(b.radius*2, b.radius*2)
	fRadius := float64(b.radius)
	dc.DrawCircle(fRadius, fRadius, fRadius)
	dc.SetRGB(1.0, 0, 0)
	dc.Fill()
	ball := ebiten.NewImageFromImage(dc.Image())
	return ball
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	screen.DrawImage(g.ground, g.groundOptions)
	g.ball.applyGravity()
	screen.DrawImage(g.ball.image, g.ball.imageOptions)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

const groundHeight = 35

func main() {

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Bouncing Ball")

	game := &Game{
		ball:          ball{radius: 20, vSpeed: 100, hSpeed: -10, vPos: 80, hPos: 10},
		ground:        ebiten.NewImage(320, groundHeight),
		groundOptions: &ebiten.DrawImageOptions{},
	}
	game.ground.Fill(color.RGBA{25, 150, 25, 0xff})
	game.groundOptions.GeoM.Translate(0.0, 240.0-groundHeight)
	game.ball.image = game.ball.makeImage()
	game.ball.imageOptions = &ebiten.DrawImageOptions{}
	game.ball.imageOptions.GeoM.Translate(float64(game.ball.hPos), float64(game.ball.vPos))

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
