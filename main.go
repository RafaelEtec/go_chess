package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_WIDTH  = 440
	SCREEN_HEIGHT = 400

	BOARD_WIDTH  = SCREEN_WIDTH - 40
	BOARD_HEIGHT = SCREEN_HEIGHT

	FRAME_OX    = 0
	FRAME_OY    = 0
	FRAME_WIDTH = 32

	COLS = 8
	ROWS = 8

	MESSAGE_WHOPLAYS = "Player %d turn"
)

type Game struct {
	state   int
	message string

	whoplay int

	board *Board
}

type Board struct {
	Pieces [][]*Piece
}

type Piece interface {
	Img() *ebiten.Image

	Color() int

	X() int
	Y() int

	C() int
	R() int
	PC() int
	PR() int
}

type Pawn struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("CHESS by Rafael Goulart")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
