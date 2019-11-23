package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	c "github.com/sylba2050/ebitest/images/chara/chara1"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

func init() {
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	DrawImageWithPosition(screen, c.Chara1_1, Position{X: 0, Y: 0})

	DrawImageWithPosition(screen, c.Chara1_2, Position{X: 40, Y: 0})

	DrawImageWithPosition(screen, c.Chara1_3, Position{X: 80, Y: 0})

	DrawImageWithPosition(screen, c.Chara1_4, Position{X: 120, Y: 0})

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Paint (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
