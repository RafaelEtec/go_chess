package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	SCREEN_WIDTH  = 256
	SCREEN_HEIGHT = 280

	FRAME_OX     = 0
	FRAME_OY     = 0
	FRAME_WIDTH  = 32
	FRAME_HEIGHT = 32
	W            = 32

	BOARD_WIDTH  = 256
	BOARD_HEIGHT = 256

	COLS = 8
	ROWS = 8

	MESSAGE_WHOPLAYS = "Player %d to move"
)

type Game struct {
	board *Board
}

type Board struct {
	pieces []*Piece
}

type Piece struct {
	img *ebiten.Image

	pieceType string
	color     string

	X int
	Y int

	C int
	R int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{180, 180, 180, 255})

	opts := ebiten.DrawImageOptions{}

	drawBoard(g, opts, screen)
	drawPieces(g, opts, screen)
}

func drawPieces(g *Game, opts ebiten.DrawImageOptions, screen *ebiten.Image) {

}

func drawBoard(g *Game, opts ebiten.DrawImageOptions, screen *ebiten.Image) {
	light_tile, _, err := ebitenutil.NewImageFromFile("assets/board/light_tile.png")
	if err != nil {
		log.Fatal(err)
	}

	dark_tile, _, err := ebitenutil.NewImageFromFile("assets/board/dark_tile.png")
	if err != nil {
		log.Fatal(err)
	}

	tileC := 0
	tile := light_tile
	for c := 0; c < COLS; c++ {
		for r := 0; r < ROWS; r++ {
			if tileC%2 == 0 {
				tile = light_tile
			} else {
				tile = dark_tile
			}

			fox, foy, fw, fh := FRAME_OX, FRAME_OY, FRAME_WIDTH, FRAME_HEIGHT
			foy += 32 * 0
			fh *= 0 + 1

			opts.GeoM.Translate(float64(c)*W, float64(r)*W)
			screen.DrawImage(
				tile.SubImage(
					image.Rect(fox, foy, fw, fh),
				).(*ebiten.Image),
				&opts,
			)
			opts.GeoM.Reset()
			tileC++
		}
		tileC++
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	game := &Game{
		board: &Board{
			pieces: []*Piece{
				{pieceType: "br"}, {pieceType: "bn"}, {pieceType: "bb"}, {pieceType: "bq"}, {pieceType: "bk"}, {pieceType: "bb"}, {pieceType: "bn"}, {pieceType: "br"},
				{pieceType: "bp"}, {pieceType: "bp"}, {pieceType: "bp"}, {pieceType: "bp"}, {pieceType: "bp"}, {pieceType: "bp"}, {pieceType: "bp"}, {pieceType: "bp"},
				{}, {}, {}, {}, {}, {}, {}, {},
				{}, {}, {}, {}, {}, {}, {}, {},
				{}, {}, {}, {}, {}, {}, {}, {},
				{}, {}, {}, {}, {}, {}, {}, {},
				{pieceType: "wp"}, {pieceType: "wp"}, {pieceType: "wp"}, {pieceType: "wp"}, {pieceType: "wp"}, {pieceType: "wp"}, {pieceType: "wp"}, {pieceType: "wp"},
				{pieceType: "wr"}, {pieceType: "wn"}, {pieceType: "wb"}, {pieceType: "wq"}, {pieceType: "wk"}, {pieceType: "wb"}, {pieceType: "wn"}, {pieceType: "wr"},
			},
		},
	}

	createBoard(game)

	ebiten.SetWindowSize(SCREEN_WIDTH*2, SCREEN_HEIGHT*2)
	ebiten.SetWindowTitle("Chess by Rafael Goulart")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func createBoard(g *Game) {
	piece := 0
	for c := 0; c < COLS; c++ {
		for r := 0; r < ROWS; r++ {
			g.board.pieces[piece].X = c * W
			g.board.pieces[piece].Y = r * W
			g.board.pieces[piece].C = c * W
			g.board.pieces[piece].R = r * W
			piece++
		}
	}
}
