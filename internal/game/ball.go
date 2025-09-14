package game

import (
	"go-pong/internal/types"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	Pos      types.Double
	Size     types.Double
	Velocity types.Double
	NumHits  float32
}

func (b *Ball) Move(g *Game) {
	if g.IsGoal {
		return 
	}
	
	if b.Pos.Y+b.Velocity.Y <= 0 || b.Pos.Y+b.Size.Y+b.Velocity.Y >= g.GameScreen.Size.Y {
		b.Velocity.Y *= -1
	}

	if b.Pos.X <= 0 {
		g.Player1.Score += 1;
		g.IsGoal = true;
	} else if b.Pos.X >= g.GameScreen.Size.X {
		g.Player2.Score += 1;
		g.IsGoal = true;
	}

	b.PlayerCollision(b.Pos.X <= g.Player1.Pos.X+g.Player1.Size.X, &g.Player1)
	b.PlayerCollision(b.Pos.X+b.Size.X >= g.Player2.Pos.X, &g.Player2)

	b.Pos.X += b.Velocity.X
	b.Pos.Y += b.Velocity.Y
}

func (b *Ball) PlayerCollision(xCollision bool, p *Player) {
	if xCollision && b.Pos.Y >= p.Pos.Y && b.Pos.Y+b.Size.Y <= p.Pos.Y+p.Size.Y {
		b.Velocity.X *= -1
		b.NumHits += 0.2

		if b.Velocity.X < float32(0) {
			b.Velocity.X -= b.NumHits
			b.Velocity.Y -= b.NumHits
		} else {
			b.Velocity.X += b.NumHits
			b.Velocity.Y += b.NumHits
		}
	}
}

func (b *Ball) DrawCircle(screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	circle := ebiten.NewImage(int(b.Size.X), int(b.Size.Y))
	vector.DrawFilledCircle(circle, float32(b.Size.X/2), float32(b.Size.Y/2), float32(b.Size.X/2), color.White, false)
	drawOptions.GeoM.Scale(1.0, 1.0)
	drawOptions.GeoM.Translate(float64(b.Pos.X), float64(b.Pos.Y))
	screen.DrawImage(circle, drawOptions)
}
