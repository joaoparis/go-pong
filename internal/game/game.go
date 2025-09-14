package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	GameState  State
	Ball       Ball
	Player1    Player
	Player2    Player
	GameScreen GameScreen
	IsGoal     bool
	IsPaused   bool
}

func (g *Game) Start() {
	ballSizeX := 20
	ballSizeY := 20
	ballVelocity := 2
	playerBoundGap := 10
	playerSizeX := 20
	playerSizeY := 100
	playerVelocity := float32(5)

	ball := Ball{
		Pos: Double{
			X: g.GameScreen.Size.X/2 - float32(ballSizeX)/2,
			Y: g.GameScreen.Size.Y/2 - float32(ballSizeY)/2,
		},
		Size: Double{
			X: float32(ballSizeX),
			Y: float32(ballSizeY),
		},
		Velocity: Double{
			X: float32(ballVelocity),
			Y: float32(ballVelocity),
		},
		NumHits: 0,
	}

	player1 := Player{
		Pos: Double{
			X: float32(playerBoundGap),
			Y: g.GameScreen.Size.Y/2 - float32(playerSizeY)/2,
		},
		Size: Double{
			X: float32(playerSizeX),
			Y: float32(playerSizeY),
		},
		Velocity: playerVelocity,
		Color:    color.RGBA{R: 255, G: 0, B: 0, A: 255},
		Keys: Keys{
			Up:   ebiten.KeyW,
			Down: ebiten.KeyS,
		},
	}

	player2 := Player{
		Pos: Double{
			X: g.GameScreen.Size.X - float32(playerSizeX+playerBoundGap),
			Y: g.GameScreen.Size.Y/2 - float32(playerSizeY)/2,
		},
		Size: Double{
			X: float32(playerSizeX),
			Y: float32(playerSizeY),
		},
		Velocity: playerVelocity,
		Color:    color.RGBA{R: 0, G: 0, B: 255, A: 255},
		Keys: Keys{
			Up:   ebiten.KeyUp,
			Down: ebiten.KeyDown,
		},
	}

	g.Ball = ball
	g.Player1 = player1
	g.Player2 = player2
	g.IsGoal = false
	g.IsPaused = false
	g.GameState = &PlayingState{}
}

func (game *Game) isGamePaused() bool {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) && !game.IsGoal {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			game.IsPaused = !game.IsPaused
		}
	}
	if game.IsPaused {
		return true
	}
	return false
}

func (g *Game) StateUpdate(screen *ebiten.Image) {
	if g.IsGoal {
		DrawText(screen, "GOAL", g.GameScreen.Size)
		lobbyState := &LobbyState{}
		g.GameState = lobbyState.New(g.GameScreen.Size)
	} else if g.IsPaused {
		DrawText(screen, "PAUSE", g.GameScreen.Size)
	}
}
