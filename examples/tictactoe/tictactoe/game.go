package tictactoe

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600
	boardSize    = 3
)

type Game struct {
	board      *Board
	boardImage *ebiten.Image
}

func NewGame() (*Game, error) {
	g := &Game{}
	var err error
	g.board, err = NewBoard(boardSize)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		g.boardImage = ebiten.NewImage(g.board.Size())
	}
	screen.Fill(backgroundColor)
	g.board.Draw(g.boardImage)
	op := _getDrawImageOptionAtGeoMMiddle(screen, g)
	screen.DrawImage(g.boardImage, op)
}

func _getDrawImageOptionAtGeoMMiddle(screen *ebiten.Image, g *Game) (drawOp *ebiten.DrawImageOptions) {
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bw, bh := g.boardImage.Bounds().Dx(), g.boardImage.Bounds().Dy()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	return op
}
