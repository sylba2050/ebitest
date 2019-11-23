package main

import (
	"github.com/hajimehoshi/ebiten"
)

var op ebiten.DrawImageOptions

func DrawImageWithPosition(dst, src *ebiten.Image, p Position) {
	op.GeoM.Translate(float64(p.X), float64(p.Y))
	dst.DrawImage(src, &op)
	op.GeoM.Reset()
}
