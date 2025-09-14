package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

type LobbyState struct {
	HostButton Button
	JoinButton Button
}

func (ls *LobbyState) New(screenSize Double) *LobbyState {
	return &LobbyState{
		HostButton: CreateHostButton(float32(screenSize.X), float32(screenSize.Y)),
		JoinButton: CreateJoinButton(float32(screenSize.X), float32(screenSize.Y)),
	}
}

func (ls *LobbyState) Layout(outsideWidth, outsideHeight *int) (screenWidth, screenHeight int) {
	return *outsideWidth, *outsideHeight
}

func (ls *LobbyState) Draw(screen *ebiten.Image, game *Game) {
	//Draw Menu
	drawTitle(screen, game.GameScreen)
	drawButton(screen, ls.HostButton)
	drawButton(screen, ls.JoinButton)
}

func (ls *LobbyState) Update(game *Game) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		game.Start()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		ls.HostButton.IsHighlighted = !ls.HostButton.IsHighlighted
		ls.JoinButton.IsHighlighted = !ls.JoinButton.IsHighlighted
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		ls.HostButton.IsHighlighted = !ls.HostButton.IsHighlighted
		ls.JoinButton.IsHighlighted = !ls.JoinButton.IsHighlighted
	}

	return nil
}

type Button struct {
	Text            string
	IsHighlighted   bool
	Pos             Double
	Size            Double
	BackgroundColor color.Color
	TextColor       color.Color
	//function on press
}

func drawTitle(screen *ebiten.Image, gs GameScreen) {
	face := text.NewGoXFace(basicfont.Face7x13)
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(float64(7*gs.Size.X/20), float64(gs.Size.Y/4))
	opts.ColorScale.ScaleWithColor(color.RGBA{R: 255, G: 0, B: 255, A: 255})
	text.Draw(screen, "MAIN MENU", face, opts)
}

func drawButton(screen *ebiten.Image, b Button) {
	highlightedBgColor := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	notHighlightedBgColor := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	highlightedTextColor := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	notHighlightedTextColor := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	bgColor := notHighlightedBgColor
	txtColor := notHighlightedTextColor
	if b.IsHighlighted {
		bgColor = highlightedBgColor
		txtColor = highlightedTextColor
	}

	drawOptions := &ebiten.DrawImageOptions{}
	rect := ebiten.NewImage(int(b.Size.X), int(b.Size.Y))
	vector.DrawFilledRect(rect, 0, 0, b.Size.X, b.Size.Y, bgColor, true)
	drawOptions.GeoM.Scale(1.0, 1.0)
	drawOptions.GeoM.Translate(float64(b.Pos.X), float64(b.Pos.Y))
	screen.DrawImage(rect, drawOptions)

	face := text.NewGoXFace(basicfont.Face7x13)
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(float64(b.Pos.X+b.Size.X/3), float64(b.Pos.Y+b.Size.Y/2))
	opts.ColorScale.ScaleWithColor(txtColor)
	text.Draw(screen, b.Text, face, opts)
}

func CreateHostButton(winX float32, winY float32) Button {
	return Button{
		Text:          "Host a game",
		IsHighlighted: true,
		Pos: Double{
			X: winX/2 - (winX/5)/2,
			Y: (2 * winY) / 5,
		},
		Size: Double{
			X: winX / 5,
			Y: winY / 8,
		},
		BackgroundColor: color.RGBA{R: 255, G: 255, B: 255, A: 255},
		TextColor:       color.RGBA{R: 0, G: 0, B: 0, A: 255},
	}
}

func CreateJoinButton(winX float32, winY float32) Button {
	return Button{
		Text:          "Join a game",
		IsHighlighted: false,
		Pos: Double{
			X: winX/2 - (winX/5)/2,
			Y: (3 * winY) / 5,
		},
		Size: Double{
			X: winX / 5,
			Y: winY / 8,
		},
		BackgroundColor: color.RGBA{R: 0, G: 0, B: 0, A: 255},
		TextColor:       color.RGBA{R: 255, G: 255, B: 255, A: 255},
	}
}
