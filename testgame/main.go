// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"fmt"
	"github.com/EbitenPot/PotCore/internal/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
	"os"
)

var img *ebiten.Image

func init() {
	var err error
	paths, _ := os.Getwd()
	img, _, err = ebitenutil.NewImageFromFile(paths + "/testgame/res/gopher.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game types.PotGame

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(img, nil)
	msg := fmt.Sprintf(`TPS: %0.2f
FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	os.Setenv("EBITEN_GRAPHICS_LIBRARY", "auto")
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Render an image")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
