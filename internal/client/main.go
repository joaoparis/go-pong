package client

import (
	"log"

	"go-pong/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func Run(game *game.Game) {
	game.Start()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

