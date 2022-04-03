package main

import (
	"log"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game interface.
type Game struct {
	ball      ball
	groundPos int
}

type ball struct {
	radius int
	vSpeed float32
	hSpeed float32
	vPos   int
	hPos   int
}

// applyGravity updates the position and speed of the ball for a single tick
func (b *ball) applyGravity() {

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
	ground := ebiten.NewImage(320, g.groundPos)
	green := color.RGBA{0, 255, 0, 0xff}
	ground.Fill(green)
	//screen.Set(240-g.groundPos, 0, nil)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0.0, 240-5.0)
	screen.DrawImage(ground, op)

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{
		ball: ball{
			radius: 20,
			vSpeed: 100,
			hSpeed: -10,
			vPos:   80,
			hPos:   10,
		},
		groundPos: 5,
	}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Bouncing Ball")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
