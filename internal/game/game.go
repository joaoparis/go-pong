package game

import (
	"fmt"
	"go-pong/internal/text_draw"
	"go-pong/internal/types"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	BallSizeX = 20
	BallSizeY = 20
	BallVelocity = 2
	MessageTicks = 120
)

type Game struct {
	GameState  State
	Ball       Ball
	Player1    Player
	Player2    Player
	GameScreen types.GameScreen
	IsGoal     bool
	IsPaused   bool
	RespawnBall	bool
	messageTicks int
}

func (g *Game) Start() {
	playerBoundGap := 10
	playerSizeX := 20
	playerSizeY := 100
	playerVelocity := float32(5)

	ball := Ball{
		Pos: types.Double{
			X: g.GameScreen.Size.X/2 - float32(BallSizeX)/2,
			Y: g.GameScreen.Size.Y/2 - float32(BallSizeY)/2,
		},
		Size: types.Double{
			X: float32(BallSizeX),
			Y: float32(BallSizeY),
		},
		Velocity: types.Double{
			X: float32(BallVelocity),
			Y: float32(BallVelocity),
		},
		NumHits: 0,
	}

	player1 := Player{
		Pos: types.Double{
			X: float32(playerBoundGap),
			Y: g.GameScreen.Size.Y/2 - float32(playerSizeY)/2,
		},
		Size: types.Double{
			X: float32(playerSizeX),
			Y: float32(playerSizeY),
		},
		Velocity: playerVelocity,
		Color:    color.RGBA{R: 255, G: 0, B: 0, A: 255},
		Keys: types.Keys{
			Up:   ebiten.KeyW,
			Down: ebiten.KeyS,
		},
		Score: 0,
	}

	player2 := Player{
		Pos: types.Double{
			X: g.GameScreen.Size.X - float32(playerSizeX+playerBoundGap),
			Y: g.GameScreen.Size.Y/2 - float32(playerSizeY)/2,
		},
		Size: types.Double{
			X: float32(playerSizeX),
			Y: float32(playerSizeY),
		},
		Velocity: playerVelocity,
		Color:    color.RGBA{R: 0, G: 0, B: 255, A: 255},
		Keys: types.Keys{
			Up:   ebiten.KeyUp,
			Down: ebiten.KeyDown,
		},
		Score: 0,
	}

	g.Ball = ball
	g.Player1 = player1
	g.Player2 = player2
	g.IsGoal = false
	g.IsPaused = false
	g.GameState = &PlayingState{}
	g.messageTicks = MessageTicks
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
		text_draw.CenterText(screen, fmt.Sprintf("GOAL\nScore: %v : %v", g.Player1.Score, g.Player2.Score), 50);
	} else if g.IsPaused {
		text_draw.CenterText(screen, "Pause", 50);
	}
}

func (g *Game) Continue(ps *PlayingState) {
	g.IsGoal = false;
	g.RespawnBall = false;

	// todo: Reset Player Position
	g.Ball = Ball{
		Pos: types.Double{
			X: g.GameScreen.Size.X/2 - float32(BallSizeX)/2,
			Y: g.GameScreen.Size.Y/2 - float32(BallSizeY)/2,
		},
		Size: types.Double{
			X: float32(BallSizeX),
			Y: float32(BallSizeY),
		},
		Velocity: types.Double{
			X: float32(BallVelocity),
			Y: float32(BallVelocity),
		},
		NumHits: 0,
	}
	g.GameState = ps;
	g.messageTicks = MessageTicks
}