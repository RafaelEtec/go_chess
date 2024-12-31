package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	SCREEN_WIDTH  = BOARD_WIDTH
	SCREEN_HEIGHT = BOARD_HEIGHT

	FRAME_OX     = 0
	FRAME_OY     = 0
	FRAME_WIDTH  = 48
	FRAME_HEIGHT = 48
	W            = 48

	BOARD_WIDTH  = W * 8
	BOARD_HEIGHT = W * 8

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
	count := 0
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			tile := g.board.pieces[count]

			if tile.pieceType != "blank" {
				fox, foy, fw, fh := FRAME_OX, FRAME_OY, FRAME_WIDTH, FRAME_HEIGHT

				opts.GeoM.Translate(float64(c)*W+3, float64(r)*W+6)
				screen.DrawImage(
					tile.img.SubImage(
						image.Rect(fox, foy, fw, fh),
					).(*ebiten.Image),
					&opts,
				)
				opts.GeoM.Reset()
			}
			count++
		}
	}
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
				{pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"},
				{pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"},
				{pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"},
				{pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"}, {pieceType: "blank"},
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
	bp, _, err := ebitenutil.NewImageFromFile("assets/pieces/black_pawn.png")
	if err != nil {
		log.Fatal(err)
	}
	wp, _, err := ebitenutil.NewImageFromFile("assets/pieces/white_pawn.png")
	if err != nil {
		log.Fatal(err)
	}

	br, _, err := ebitenutil.NewImageFromFile("assets/pieces/black_rook.png")
	if err != nil {
		log.Fatal(err)
	}
	wr, _, err := ebitenutil.NewImageFromFile("assets/pieces/white_rook.png")
	if err != nil {
		log.Fatal(err)
	}

	bb, _, err := ebitenutil.NewImageFromFile("assets/pieces/black_bishop.png")
	if err != nil {
		log.Fatal(err)
	}
	wb, _, err := ebitenutil.NewImageFromFile("assets/pieces/white_bishop.png")
	if err != nil {
		log.Fatal(err)
	}

	bn, _, err := ebitenutil.NewImageFromFile("assets/pieces/black_knight.png")
	if err != nil {
		log.Fatal(err)
	}
	wn, _, err := ebitenutil.NewImageFromFile("assets/pieces/white_knight.png")
	if err != nil {
		log.Fatal(err)
	}

	bq, _, err := ebitenutil.NewImageFromFile("assets/pieces/black_queen.png")
	if err != nil {
		log.Fatal(err)
	}
	wq, _, err := ebitenutil.NewImageFromFile("assets/pieces/white_queen.png")
	if err != nil {
		log.Fatal(err)
	}

	bk, _, err := ebitenutil.NewImageFromFile("assets/pieces/black_king.png")
	if err != nil {
		log.Fatal(err)
	}
	wk, _, err := ebitenutil.NewImageFromFile("assets/pieces/white_king.png")
	if err != nil {
		log.Fatal(err)
	}

	piece := 0
	for c := 0; c < COLS; c++ {
		for r := 0; r < ROWS; r++ {
			tile := g.board.pieces[piece]
			switch tile.pieceType {
			case "bp":
				tile.img = bp
			case "wp":
				tile.img = wp
			case "br":
				tile.img = br
			case "wr":
				tile.img = wr
			case "bb":
				tile.img = bb
			case "wb":
				tile.img = wb
			case "bn":
				tile.img = bn
			case "wn":
				tile.img = wn
			case "bq":
				tile.img = bq
			case "wq":
				tile.img = wq
			case "bk":
				tile.img = bk
			case "wk":
				tile.img = wk
			}

			tile.X = c * W
			tile.Y = r * W
			tile.C = c
			tile.R = r
			piece++
		}
	}
}
