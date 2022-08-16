package game

import (
	"errors"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/patricktcoakley/go-rtiow/color"
)

var errShutdown = errors.New("Shutdown")

type Game struct {
	width, height int
	buffer        []byte
}

func New(width, height int, title string) *Game {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)
	return &Game{width, height, make([]byte, width*height*4)}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errShutdown
	}
	return nil
}

func (g *Game) WritePixel(x, y int, color *color.Color) {
	y = g.height - y - 1
	offset := 4 * (y*g.width + x)
	g.buffer[offset] = color.R
	g.buffer[offset+1] = color.G
	g.buffer[offset+2] = color.B
	g.buffer[offset+3] = 255
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(g.buffer)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}

func (g *Game) Run() {
	if err := ebiten.RunGame(g); err != nil {
		if err == errShutdown {
			os.Exit(0)
		}
		log.Fatal(err)
	}
}
