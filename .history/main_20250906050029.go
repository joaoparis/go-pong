package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/basicfont"
)

type Double struct {
	X float32
	Y float32
}

type Keys struct {
	Up ebiten.Key
	Down ebiten.Key
}

type Player struct {
	Pos  Double
	Size Double
	Velocity float32
	Color color.Color
	Keys Keys
}

func (p *Player) MoveY(screen *GameScreen) {
	if ebiten.IsKeyPressed(p.Keys.Up) && (p.Pos.Y - p.Velocity) >= 0 {
		p.Pos.Y -= p.Velocity
	}

	if ebiten.IsKeyPressed(p.Keys.Down) && (p.Pos.Y + p.Velocity + p.Size.Y) <= screen.Size.Y {
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

type Ball struct {
	Pos  Double
	Size Double
	Velocity Double
	NumHits float32
}

func (b *Ball) Move(g *Game) {
	if b.Pos.Y + b.Velocity.Y <= 0 || b.Pos.Y + b.Size.Y + b.Velocity.Y >= g.GameScreen.Size.Y {
		b.Velocity.Y*=-1
	}
	if b.Pos.X + b.Velocity.X <= 0 || b.Pos.X + b.Size.X + b.Velocity.X == g.GameScreen.Size.X {
		print("GOAL!!!!")
		g.isGoal = true
	}
	
	b.PlayerCollision(b.Pos.X <= g.Player1.Pos.X + g.Player1.Size.X, &g.Player1)
	b.PlayerCollision(b.Pos.X + b.Size.X >= g.Player2.Pos.X, &g.Player2)
	
	b.Pos.X += b.Velocity.X
	b.Pos.Y += b.Velocity.Y
}

func (b *Ball) PlayerCollision(xCollision bool, p *Player) {
	if xCollision && b.Pos.Y >= p.Pos.Y && b.Pos.Y + b.Size.Y <= p.Pos.Y + p.Size.Y {
		b.Velocity.X*=-1
		b.NumHits+=0.2
		
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

	drawOptions.GeoM.Scale(1.0, 1.0)                               // stretch X by 2 → ellipse
	drawOptions.GeoM.Translate(float64(b.Pos.X), float64(b.Pos.Y)) // place at (0,0)
	screen.DrawImage(circle, drawOptions)
}
 
type GameScreen struct {
	Size Double
}

type Game struct {
	Ball    Ball
	Player1 Player
	Player2 Player
	GameScreen GameScreen
	isGoal bool
	isPaused bool
}

func (game *Game) Update() error {

	//is this giving priority to player1? 
	//could this be done in sync?
	game.Ball.Move(game)
	game.Player1.MoveY(&game.GameScreen)
	game.Player2.MoveY(&game.GameScreen)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Ball.DrawCircle(screen)
	g.Player1.DrawRect(screen)
	g.Player2.DrawRect(screen)
	g.StateUpdate(screen)
}

func (g *Game) StateUpdate(screen *ebiten.Image) {

	if g.isGoal {
		drawOptions := &ebiten.DrawImageOptions{}

		rect := ebiten.NewImage(int(g.GameScreen.Size.X/2), int(g.GameScreen.Size.Y/2))
		vector.DrawFilledRect(rect, 0, 0, g.GameScreen.Size.X/2, g.GameScreen.Size.Y/2, color.White, true)

		drawOptions.GeoM.Scale(1.0, 1.0)
		drawOptions.GeoM.Translate(float64(g.GameScreen.Size.X/4), float64(g.GameScreen.Size.Y/4))
		screen.DrawImage(rect, drawOptions)

		DrawText(screen, "GOAL!")
		DrawText(screen, "GOAL!")
		DrawText(screen, "GOAL!")
	} else if g.isPaused {
        face := text.NewGoXFace(basicfont.Face7x13)
        opts := &text.DrawOptions{}
        opts.GeoM.Translate(float64(g.GameScreen.Size.X/2-40), float64(g.GameScreen.Size.Y/2))
        opts.ColorScale.ScaleWithColor(color.RGBA{R: 255, G: 255, B: 255, A: 255})
        text.Draw(screen, "PAUSED", face, opts)
    }
}

func DrawText(screen *ebiten.Image, text String) void {
		face := text.NewGoXFace(basicfont.Face7x13)
        opts := &text.DrawOptions{}
  		opts.GeoM.Translate(float64(g.GameScreen.Size.X/4), float64(g.GameScreen.Size.Y/4))
        opts.ColorScale.ScaleWithColor(color.RGBA{R: 255, G: 0, B: 0, A: 255})
    	text.Draw(screen, text, face, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	winX := 800
	winY := 480
	ebiten.SetWindowSize(winX, winY)
	ebiten.SetWindowTitle("Pong")
	ebiten.SetWindowPosition(10, 10)

	ballSizeX := 20
	ballSizeY := 20
	ballVelocity := float32(2)
	playerBoundGap := 5
	playerSizeX := 20
	playerSizeY := 100
	playerVelocity := float32(5)

	gameScreen := &GameScreen{
		Size: Double{
			X: float32(winX),
			Y: float32(winY),
		},
	}
	
	ball := &Ball{
		Pos: Double{
			X: float32(winX/2 - ballSizeX/2),
			Y: float32(winY/2 - ballSizeY/2),
		},
		Size: Double{
			X: float32(ballSizeX),
			Y: float32(ballSizeY),
		},
		Velocity: Double{
			X: ballVelocity,
			Y: ballVelocity,
		},
		NumHits: 0,
	}

	player1 := &Player{
		Pos: Double{
			X: float32(playerBoundGap),
			Y: float32(winY/2 - playerSizeY/2),
		},
		Size: Double{
			X: float32(playerSizeX),
			Y: float32(playerSizeY),
		},
		Velocity: playerVelocity,
		Color: color.RGBA{R: 255, G: 0, B: 0, A: 255},
		Keys: Keys{
			Up: ebiten.KeyW,
			Down: ebiten.KeyS,
		},
	}

	player2 := &Player{
		Pos: Double{
			X: float32(winX - playerSizeX - playerBoundGap),
			Y: float32(winY/2 - playerSizeY/2),
		},
		Size: Double{
			X: float32(playerSizeX),
			Y: float32(playerSizeY),
		},
		Velocity: playerVelocity,
		Color: color.RGBA{R: 0, G: 0, B: 255, A: 255},
		Keys: Keys{
			Up: ebiten.KeyUp,
			Down: ebiten.KeyDown,
		},
	}

	g := &Game{
		Ball: *ball,
		Player1: *player1,
		Player2: *player2,
		GameScreen: *gameScreen,
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
