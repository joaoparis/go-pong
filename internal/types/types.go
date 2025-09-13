package types

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
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

type Ball struct {
	Pos  Double
	Size Double
	Velocity Double
	NumHits float32
}

type GameScreen struct {
	Size Double
}

type Game struct {
	Window Double
	Ball    Ball
	Player1 Player
	Player2 Player
	GameScreen GameScreen
	IsGoal bool
	IsPaused bool
}