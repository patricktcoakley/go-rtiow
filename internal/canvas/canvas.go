package canvas

import (
	"errors"
	"image/color"
	"log"
	"math"
	"os"

	"github.com/patricktcoakley/go-rtiow/internal/vec3"

	"github.com/hajimehoshi/ebiten/v2"
)

var errShutdown = errors.New("Shutdown")

type Canvas struct {
	width, height   int
	samplesPerPixel float64
	buffer          []byte
}

func NewCanvas(width, height, samplesPerPixel int, title string) *Canvas {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)
	return &Canvas{
		width,
		height,
		float64(samplesPerPixel),
		make([]byte, width*height*4)}
}

func (c *Canvas) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errShutdown
	}
	return nil
}

func (c *Canvas) pixelOffset(x, y int) int {
	y = c.height - y - 1
	return 4 * (y*c.width + x)
}

func (c *Canvas) WritePixel(x, y int, color vec3.Vec3) {
	scale := 1 / c.samplesPerPixel
	pixelColor := newColorFromVec3(color, scale)
	offset := c.pixelOffset(x, y)
	c.buffer[offset] = pixelColor.R
	c.buffer[offset+1] = pixelColor.G
	c.buffer[offset+2] = pixelColor.B
	c.buffer[offset+3] = pixelColor.A
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

func newColorFromVec3(v vec3.Vec3, scale float64) color.RGBA {
	return color.RGBA{
		uint8(clamp(math.Sqrt(v[0]*scale)) * 256),
		uint8(clamp(math.Sqrt(v[1]*scale)) * 256),
		uint8(clamp(math.Sqrt(v[2]*scale)) * 256),
		255,
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
