package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type PlayingState struct {
	LPlayerScore int
	RPlayerScore int
}

func (ps *PlayingState) Layout(outsideWidth, outsideHeight *int) (screenWidth, screenHeight int) {
	return *outsideWidth, *outsideHeight
}

func (ps *PlayingState) Draw(screen *ebiten.Image, game *Game) {
	game.Ball.DrawCircle(screen)
	game.Player1.DrawRect(screen)
	game.Player2.DrawRect(screen)
	game.StateUpdate(screen)
}

func (ps *PlayingState) Update(game *Game) error {
	// if ebiten.IsKeyPressed(ebiten.KeyR) {
	// 	game.Start()
	// }
	if game.isGamePaused() {
		return nil
	}
	if game.IsGoal {
		game.messageTicks--

		if game.messageTicks == 0 {
			game.RespawnBall = true;
		}
	}
	if game.RespawnBall {
		game.Continue(ps)
	}

	game.Ball.Move(game)
	game.Player1.MoveY(&game.GameScreen)
	game.Player2.MoveY(&game.GameScreen)

	return nil
}
