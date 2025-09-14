package game

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.GameState.Draw(screen, game)
}

func (game *Game) Update() error {
	return game.GameState.Update(game)
}

type State interface {
	Update(*Game) error
	Draw(*ebiten.Image, *Game)
	Layout(*int, *int) (int, int)
}
