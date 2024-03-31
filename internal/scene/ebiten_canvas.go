package scene

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/patricktcoakley/go-rtiow/internal/geometry"
	"github.com/patricktcoakley/go-rtiow/internal/math"
	"log"
	"os"
)

var errShutdown = errors.New("Shutdown")

type EbitenCanvas struct {
	width, height   int
	samplesPerPixel math.Real
	buffer          []byte
}

func newEbitenCanvas(width, height, samplesPerPixel int) Canvas {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Raytracing In One Weekend With Go!")
	return &EbitenCanvas{
		width,
		height,
		math.Real(samplesPerPixel),
		make([]byte, width*height*4),
	}
}

func (c *EbitenCanvas) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errShutdown
	}
	return nil
}

func (c *EbitenCanvas) pixelOffset(x, y int) int {
	y = c.height - y - 1
	return 4 * (y*c.width + x)
}

func (c *EbitenCanvas) WritePixel(x, y int, color geometry.Vec3) {
	scale := 1 / c.samplesPerPixel
	pixelColor := newRGBAFromVec3(color, scale)
	offset := c.pixelOffset(x, y)
	c.buffer[offset] = pixelColor.R
	c.buffer[offset+1] = pixelColor.G
	c.buffer[offset+2] = pixelColor.B
	c.buffer[offset+3] = pixelColor.A
}

func (c *EbitenCanvas) Draw(screen *ebiten.Image) {
	screen.WritePixels(c.buffer)
}

func (c *EbitenCanvas) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return c.width, c.height
}

func (c *EbitenCanvas) Run() {
	if err := ebiten.RunGame(c); err != nil {
		if errors.Is(err, errShutdown) {
			os.Exit(0)
		}
		log.Fatal(err)
	}
}
