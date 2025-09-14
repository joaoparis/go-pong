package game

import (
	"go-pong/internal/text_draw"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type LobbyState struct {
	Menu text_draw.Menu
}

func (ls *LobbyState) Layout(outsideWidth, outsideHeight *int) (screenWidth, screenHeight int) {
	return *outsideWidth, *outsideHeight
}

func (ls *LobbyState) Draw(screen *ebiten.Image, game *Game) {
	text_draw.DrawMenu(ls.Menu, screen, game.GameScreen);
}

func (ls *LobbyState) Update(game *Game) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		ls.Menu.Items[ls.Menu.Select].Action(game);
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		isFirst := ls.Menu.Select == 0;
		if isFirst {
			ls.Menu.Select = 0;
		} else {
			ls.Menu.Select -= 1;
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		len := len(ls.Menu.Items) - 1;
		isLast := ls.Menu.Select == len;
		if isLast {
			ls.Menu.Select = len;
		} else {
			ls.Menu.Select += 1;
		}
	}

	return nil
}