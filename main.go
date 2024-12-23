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

	BOARD_WIDTH  = 400
	BOARD_HEIGHT = 400

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

type Piece struct {
	Img *ebiten.Image

	Color int

	X int
	Y int

	C  int
	R  int
	PC int
	PR int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	opts := ebiten.DrawImageOptions{}

	drawBoard(opts, screen)
}

func drawBoard(opts ebiten.DrawImageOptions, screen *ebiten.Image) {
	light_tile, err := ebitenutil.NewImageFromURL("https://github.com/RafaelEtec/go_chess/blob/4ed5aef3d0bfe7485eb1bd0069a24ba9c528773b/assets/board/light_tile.png?raw=true")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < COLS; i++ {
		for j := 0; j < ROWS; j++ {
			opts.GeoM.Translate(float64(i)*FRAME_WIDTH, float64(j)*FRAME_WIDTH)
			screen.DrawImage(
				light_tile.SubImage(
					image.Rect(32, 32, 0, 0),
				).(*ebiten.Image),
				&opts,
			)
			opts.GeoM.Reset()
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
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
