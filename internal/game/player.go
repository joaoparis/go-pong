package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	Pos      Double
	Size     Double
	Velocity float32
	Color    color.Color
	Keys     Keys
}

func (p *Player) MoveY(screen *GameScreen) {
	if ebiten.IsKeyPressed(p.Keys.Up) && (p.Pos.Y-p.Velocity) >= 0 {
		p.Pos.Y -= p.Velocity
	}
	if ebiten.IsKeyPressed(p.Keys.Down) && (p.Pos.Y+p.Velocity+p.Size.Y) <= screen.Size.Y {
		p.Pos.Y += p.Velocity
	}
}

func (p *Player) DrawRect(screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	rect := ebiten.NewImage(int(p.Size.X), int(p.Size.Y))
	vector.DrawFilledRect(rect, 0, 0, p.Size.X, p.Size.Y, p.Color, true)
	drawOptions.GeoM.Scale(1.0, 1.0)
	drawOptions.GeoM.Translate(float64(p.Pos.X), float64(p.Pos.Y))
	screen.DrawImage(rect, drawOptions)
}
