package scene

import (
	"image"
	"image/png"
	"os"

	"github.com/patricktcoakley/go-rtiow/internal/math"
)

type PngCanvas struct {
	width, height   int
	samplesPerPixel math.Real
	img             *image.RGBA
}

func newPngCanvas(width, height, samplesPerPixel int) Canvas {
	return &PngCanvas{
		width,
		height,
		math.Real(samplesPerPixel),
		image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

func (c *PngCanvas) pixelOffset(x, y int) int {
	y = c.height - y - 1
	return 4 * (y*c.width + x)
}

func (c *PngCanvas) WritePixel(pixel Pixel) {
	scale := 1 / c.samplesPerPixel
	pixelColor := newRGBAFromVec3(pixel.Color, scale)
	offset := c.pixelOffset(pixel.X, pixel.Y)
	c.img.Pix[offset] = pixelColor.R
	c.img.Pix[offset+1] = pixelColor.G
	c.img.Pix[offset+2] = pixelColor.B
	c.img.Pix[offset+3] = pixelColor.A
}

func (c *PngCanvas) Run() {
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
