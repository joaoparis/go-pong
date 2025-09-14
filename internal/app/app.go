package app

import (
	"go-pong/internal/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func Start() {
	winX := 800
	winY := 480
	ebiten.SetWindowSize(winX, winY)
	ebiten.SetWindowTitle("Pong")
	ebiten.SetWindowPosition(10, 10)

	lobby := game.LobbyState{}

	game := &game.Game{
		GameState: lobby.New(game.Double{
			X: float32(winX),
			Y: float32(winY),
		},
		),
		GameScreen: game.GameScreen{
			Size: game.Double{
				X: float32(winX),
				Y: float32(winY),
			},
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
