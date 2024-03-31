package scene

import (
	"image"
	"image/png"
	"os"

	"github.com/patricktcoakley/go-rtiow/internal/geometry"
	"github.com/patricktcoakley/go-rtiow/internal/math"
)

type PpmCanvas struct {
	width, height   int
	samplesPerPixel math.Real
	img             *image.RGBA
}

func newPpmCanvas(width, height, samplesPerPixel int) Canvas {
	return &PpmCanvas{
		width,
		height,
		math.Real(samplesPerPixel),
		image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

func (c *PpmCanvas) pixelOffset(x, y int) int {
	y = c.height - y - 1
	return 4 * (y*c.width + x)
}

func (c *PpmCanvas) WritePixel(x, y int, color geometry.Vec3) {
	scale := 1 / c.samplesPerPixel
	pixelColor := newRGBAFromVec3(color, scale)
	offset := c.pixelOffset(x, y)
	c.img.Pix[offset] = pixelColor.R
	c.img.Pix[offset+1] = pixelColor.G
	c.img.Pix[offset+2] = pixelColor.B
	c.img.Pix[offset+3] = pixelColor.A
}

func (c *PpmCanvas) Run() {
	f, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}

	err = png.Encode(f, c.img)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	return
}
