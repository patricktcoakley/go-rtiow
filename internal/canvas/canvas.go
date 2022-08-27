package canvas

import (
	"errors"
	"log"
	"os"

	"github.com/patricktcoakley/go-rtiow/internal/vec3"

	"github.com/hajimehoshi/ebiten/v2"
)

var errShutdown = errors.New("Shutdown")

type RGB [3]uint8

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

func (c *Canvas) WritePixel(x, y int, color vec3.Vec3) {
	pixelColor := newColorFromVec3(color)
	y = c.height - y - 1
	offset := 4 * (y*c.width + x)
	c.buffer[offset] = pixelColor[0]
	c.buffer[offset+1] = pixelColor[1]
	c.buffer[offset+2] = pixelColor[2]
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

func newColorFromVec3(v vec3.Vec3) RGB {
	return RGB{
		uint8(clamp(v[0]) * 256),
		uint8(clamp(v[1]) * 256),
		uint8(clamp(v[2]) * 256),
	}
}

func clamp(x float64) float64 {
	if x < 0.0 {
		return 0.0
	}
	if x > 0.999 {
		return 0.999
	}

	return x
}
