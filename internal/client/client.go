package client

import "go-pong/internal/game"

func Run(game *game.Game) {
	game.Start()
}
