package scene

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/patricktcoakley/go-rtiow/internal/geometry"
	"github.com/patricktcoakley/go-rtiow/internal/math"
)

type Canvas struct {
	width, height   int
	samplesPerPixel math.Real
	img             *image.RGBA
}

func NewCanvas(width, height, samplesPerPixel int) Canvas {
	return Canvas{
		width,
		height,
		math.Real(samplesPerPixel),
		image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

func (c Canvas) pixelOffset(x, y int) int {
	y = c.height - y - 1
	return 4 * (y*c.width + x)
}

func (c Canvas) WritePixel(x, y int, color geometry.Vec3) {
	scale := 1 / c.samplesPerPixel
	pixelColor := newColorFromVec3(color, scale)
	offset := c.pixelOffset(x, y)
	c.img.Pix[offset] = pixelColor.R
	c.img.Pix[offset+1] = pixelColor.G
	c.img.Pix[offset+2] = pixelColor.B
	c.img.Pix[offset+3] = pixelColor.A
}

func (c Canvas) WriteImage() {
	f, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, c.img)
	if err != nil {
		panic(err)
	}
}

func newColorFromVec3(v geometry.Vec3, scale math.Real) color.RGBA {
	return color.RGBA{
		R: uint8(clamp(math.Sqrt(v.X*scale)) * 256),
		G: uint8(clamp(math.Sqrt(v.Y*scale)) * 256),
		B: uint8(clamp(math.Sqrt(v.Z*scale)) * 256),
		A: 255,
	}
}

func clamp(x math.Real) math.Real {
	if x < 0.0 {
		return 0.0
	}

	if x > 0.999 {
		return 0.999
	}

	return x
}
