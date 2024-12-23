package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_WIDTH  = 440
	SCREEN_HEIGHT = 400

	BOARD_WIDTH  = SCREEN_WIDTH - 40
	BOARD_HEIGHT = SCREEN_HEIGHT

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
	Pieces []*Piece
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

type Blank struct{}

type Pawn struct{}

type Knight struct{}

type Bishop struct{}

type Queen struct{}

type King struct{}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	opts := ebiten.DrawImageOptions{}

	drawBoard(g, opts, screen)
}

func drawBoard(g *Game, opts ebiten.DrawImageOptions, screen *ebiten.Image) {
	light_tile, err := ebitenutil.NewImageFromURL("assets/board/light_tile.png")
	if err != nil {
		log.Fatal(err)
	}

	dark_tile, err := ebitenutil.NewImageFromURL("assets/board/dark_tile.png")
	if err != nil {
		log.Fatal(err)
	}

	tile := light_tile
	for i := 0; i < COLS; i++ {
		color := 0
		for j := 0; j < ROWS; j++ {
			if color%2 == 0 {
				tile = light_tile
				color = 2
			} else {
				tile = dark_tile
				color = 0
			}

			opts.GeoM.Translate(float64(i)*FRAME_WIDTH, float64(j)*FRAME_WIDTH)
			screen.DrawImage(
				tile.SubImage(
					image.Rect(32, 32, 32, 32),
				).(*ebiten.Image),
				&opts,
			)

			opts.GeoM.Reset()
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	g := &Game{}

	ebiten.SetWindowSize(SCREEN_WIDTH*2, SCREEN_HEIGHT*2)
	ebiten.SetWindowTitle("CHESS by Rafael Goulart")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
