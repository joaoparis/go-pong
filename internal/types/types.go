package types

import "github.com/hajimehoshi/ebiten/v2"

type Double struct {
	X float32
	Y float32
}

type Keys struct {
	Up   ebiten.Key
	Down ebiten.Key
}

type GameScreen struct {
	Size Double
}
