package main

import (
	"go-pong/internal/client"

	"go-pong/internal/game"
	"go-pong/internal/types"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	winX := 800
	winY := 480
	ebiten.SetWindowSize(winX, winY)
	ebiten.SetWindowTitle("Pong")
	ebiten.SetWindowPosition(10, 10)

	game := &game.Game{
		Window: types.Double {
			X: float32(winX),
			Y: float32(winY),
		},
	}

	client.Run(game);
}
