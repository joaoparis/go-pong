package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/basicfont"
)

type PlayingState struct {
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
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		game.Start()
	}
	if game.isGamePaused() {
		return nil
	}
	if game.IsGoal {
		return nil
	}

	game.Ball.Move(game)
	game.Player1.MoveY(&game.GameScreen)
	game.Player2.MoveY(&game.GameScreen)

	return nil
}

func DrawText(screen *ebiten.Image, msg string, size Double) {
	face := text.NewGoXFace(basicfont.Face7x13)
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(float64(size.X/2), float64(size.Y/2))
	opts.ColorScale.ScaleWithColor(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	text.Draw(screen, msg, face, opts)
}
