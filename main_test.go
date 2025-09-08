package main

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

// type MockDrawer struct {
// 	Calls []string
// }

// func (m *MockDrawer) DrawText(screen *ebiten.Image, s string, x, y int) {
// 	m.Calls = append(m.Calls, s)
// }

// func TestDrawBall(t *testing.T) {
// 	mock := &MockDrawer{}
// 	game := &Game{Drawer: mock}

// 	screen := ebiten.NewImage(50, 50)
// 	game.Draw(screen)

// 	if len(mock.Calls) == 0 {
// 		t.Fatal("expected DrawText to be called, but it was not")
// 	}

// 	if mock.Calls[0] != "O" {
// 		t.Fatalf("expected 'O' to be drawn, got %s", mock.Calls[0])
// 	}
// }

// func TestBallMovesDown(t *testing.T) {
// 	mock := &MockDrawer{}
// 	game := &Game{Drawer: mock}

// 	screen := ebiten.NewImage(50, 50)
// 	game.Draw(screen)

// 	if mock.Calls[0] != "O" {
// 		t.Fatalf("expected 'O' to be drawn, got %s", mock.Calls[0])
// 	}

// 	// Simulate 5 frames
// 	for i := 0; i < 5; i++ {
// 		game.Update()
// 	}

// 	if mock.Calls[0] != "O" {
// 		t.Fatalf("expected 'O' to be drawn, got %s", mock.Calls[0])
// 	}

// }

func TestDrawBallPixels(t *testing.T) {
	screen := ebiten.NewImage(50, 50)
	DrawBall(screen, 10, 20)

	found := false
	for y := 15; y < 25; y++ {
		for x := 5; x < 15; x++ {
			_, _, _, a := screen.At(x, y).RGBA()
			if a > 0 {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		t.Fatal("expected non-transparent pixels where the ball should be drawn")
	}
}
