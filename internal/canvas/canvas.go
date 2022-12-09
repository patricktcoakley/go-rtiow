package canvas

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

type Canvas struct {
	width, height   int
	samplesPerPixel float64
	img             *image.RGBA
}

func NewCanvas(width, height, samplesPerPixel int) *Canvas {
	return &Canvas{
		width,
		height,
		float64(samplesPerPixel),
		image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

func (c Canvas) pixelOffset(x, y int) int {
	y = c.height - y - 1
	return 4 * (y*c.width + x)
}

func (c Canvas) WritePixel(x, y int, color vec3.Vec3) {
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
	png.Encode(f, c.img)
}

func newColorFromVec3(v vec3.Vec3, scale float64) color.RGBA {
	return color.RGBA{
		uint8(clamp(math.Sqrt(v.X*scale)) * 256),
		uint8(clamp(math.Sqrt(v.Y*scale)) * 256),
		uint8(clamp(math.Sqrt(v.Z*scale)) * 256),
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
