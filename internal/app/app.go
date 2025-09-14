package app

import (
	"fmt"
	"go-pong/internal/game"
	"go-pong/internal/text_draw"
	"go-pong/internal/types"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func Start() {
	winX := 800
	winY := 480
	ebiten.SetWindowSize(winX, winY)
	ebiten.SetWindowTitle("Pong")
	ebiten.SetWindowPosition(10, 10)

	lobby := game.LobbyState{
		Menu: text_draw.Menu{
			Items: []text_draw.MenuItem{
				{ Name: "Local Game", Action: func(game text_draw.GameContext) {
					game.Start();
				} },
				{ Name: "Host Game", Action: func(game text_draw.GameContext) {
					fmt.Println("Host game")
				} },
				{ Name: "Join Game", Action: func(game text_draw.GameContext) {
					fmt.Println("Join game")
				} },
				{ Name: "Quit Game", Action: func(game text_draw.GameContext) {
					fmt.Println("Quit game")
				} },
			},
			Select: 0,
			Title: "Pong",
		},
	}

	game := &game.Game{
		GameState: &lobby,
		GameScreen: types.GameScreen{
			Size: types.Double{
				X: float32(winX),
				Y: float32(winY),
			},
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
