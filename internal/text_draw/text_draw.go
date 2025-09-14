package text_draw

import (
	_ "embed"
	"go-pong/internal/types"
	"image/color"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/font/opentype"
)

type MenuItem struct {
	Name string
	Action func()
}

type Menu struct {
	Items 	[]MenuItem
	Select 	int
	Title	string
}

func CenterText(screen *ebiten.Image, msg string) {
	face := text.NewGoXFace(basicfont.Face7x13)
	screenWidth, screenHeight := screen.Size()

	// Split text into lines
	lines := strings.Split(msg, "\n")

	// Measure height of one line (approximate line spacing factor = 1.2)
	_, lineHeight := text.Measure("M", face, 1)
	totalHeight := int(float64(lineHeight) * 1.2 * float64(len(lines)))

	// Start vertical position so all lines are vertically centered
	startY := (screenHeight - totalHeight) / 2

	for i, line := range lines {
		// Measure line width
		lineWidth, _ := text.Measure(line, face, 1)

		// Horizontal center
		x := (screenWidth - int(lineWidth)) / 2

		// Vertical position for this line
		y := startY + int(float64(i)*float64(lineHeight)*1.2)

		options := &text.DrawOptions{
			LayoutOptions: text.LayoutOptions{
				PrimaryAlign:   text.AlignCenter,
				SecondaryAlign: text.AlignStart,
			},
		}
		options.GeoM.Translate(float64(x), float64(y))
		text.Draw(screen, line, face, options)
	}
}

func drawCenteredText(screen *ebiten.Image, msg string, face text.Face, screenWidth, y int, col color.Color) {
	width, _ := text.Measure(msg, face, 1)
	x := (screenWidth - int(width)) / 2

	opts := &text.DrawOptions{}
	opts.GeoM.Translate(float64(x), float64(y))
	opts.ColorScale.ScaleWithColor(col)

	text.Draw(screen, msg, face, opts)
}

//go:embed assets/fonts/Corptic.otf
var robotoBoldTTF []byte

func loadFontFace(size float64) text.Face {
    ft, err := opentype.Parse(robotoBoldTTF)
    if err != nil {
        log.Fatal(err)
    }

    fnt, err := opentype.NewFace(ft, &opentype.FaceOptions{
        Size:    size,
        DPI:     72,
        Hinting: font.HintingFull,
    })
    if err != nil {
        log.Fatal(err)
    }

    return text.NewGoXFace(fnt)
}


func DrawMenu(menu Menu, screen *ebiten.Image, screenSize types.GameScreen) {
	face := loadFontFace(60)
	screenWidth, screenHeight := screen.Size()

	// Calculate base Y position
	_, lineHeight := text.Measure("M", face, 1)
	lineSpacing := int(float64(lineHeight) * 1.1)

	startY := (screenHeight - (len(menu.Items)+1)*lineSpacing) / 2

	// Draw Title (centered)
	drawCenteredText(screen, menu.Title, face, screenWidth, startY, color.RGBA{97, 18, 85, 1})

	// Draw each menu item
	for i, item := range menu.Items {
		face := loadFontFace(25)
		y := startY + (i+1) * lineSpacing

		// Highlight selected item
		var col color.Color = color.White
		if i == menu.Select {
			col = color.RGBA{255, 255, 0, 255} // Yellow for selected
		}

		drawCenteredText(screen, item.Name, face, screenWidth, y, col)
	}
}
