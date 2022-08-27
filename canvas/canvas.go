package canvas

import (
	"errors"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/patricktcoakley/go-rtiow/geom"
)

var errShutdown = errors.New("Shutdown")

type Canvas struct {
	width, height int
	buffer        []byte
}

func NewCanvas(width, height int, title string) *Canvas {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)
	return &Canvas{width, height, make([]byte, width*height*4)}
}

func (c *Canvas) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errShutdown
	}
	return nil
}

func (c *Canvas) WritePixel(x, y int, color geom.Color) {
	y = c.height - y - 1
	offset := 4 * (y*c.width + x)
	c.buffer[offset] = color[0]
	c.buffer[offset+1] = color[1]
	c.buffer[offset+2] = color[2]
	c.buffer[offset+3] = 255
}

func (c *Canvas) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(c.buffer)
}

func (c *Canvas) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return c.width, c.height
}

func (c *Canvas) Run() {
	if err := ebiten.RunGame(c); err != nil {
		if err == errShutdown {
			os.Exit(0)
		}
		log.Fatal(err)
	}
}
