package tictactoe

import (
	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
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
	circleImg  *ebiten.Image
	touchPos   *Pos
}

type Pos struct {
	x int
	y int
}

func NewBoard(size int) (*Board, error) {
	b := &Board{
		size: size,
	}
	if b.blackPixel == nil {
		b.blackPixel = ebiten.NewImage(linePixel, linePixel)
		b.blackPixel.Fill(color.Black)
	}
	if b.circleImg == nil {
		img, _, err := image.Decode(bytes.NewReader(images.Circle_jpg))
		if err != nil {
			log.Fatal(err)
		}
		b.circleImg = ebiten.NewImageFromImage(img)
	}
	if b.touchPos == nil {
		b.touchPos = &Pos{x: 83, y: 172}
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

	if b.touchPos != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(b.touchPos.x), float64(b.touchPos.y))
		op.ColorScale.ScaleWithColor(frameColor)
		boardImage.DrawImage(b.circleImg, op)
	}
}
