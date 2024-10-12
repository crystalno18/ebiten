package tictactoe

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	tileSize   = 80
	tileMargin = 4
	lineMargin = 9
	linePixel  = 3
)

type Board struct {
	size       int
	blackPixel *ebiten.Image
}

func NewBoard(size int) (*Board, error) {
	b := &Board{
		size: size,
	}
	if b.blackPixel == nil {
		b.blackPixel = ebiten.NewImage(linePixel, linePixel)
		b.blackPixel.Fill(color.Black)
	}
	return b, nil
}

func (b *Board) Size() (int, int) {
	x := b.size*tileSize + (b.size+1)*tileMargin
	y := x
	return x, y
}

func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(frameColor)
	bx, by := b.Size()
	for i := 0; i < 2; i++ {
		y := float64(by) / 3.0 * (float64(i) + 1)
		for x := lineMargin; x < bx-lineMargin; x++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			op.ColorScale.ScaleWithColor(frameColor)
			boardImage.DrawImage(b.blackPixel, op)
		}
	}
	for i := 0; i < 2; i++ {
		x := float64(bx) / 3.0 * (float64(i) + 1)
		for y := lineMargin; y < by-lineMargin; y++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			op.ColorScale.ScaleWithColor(frameColor)
			boardImage.DrawImage(b.blackPixel, op)
		}
	}

}
