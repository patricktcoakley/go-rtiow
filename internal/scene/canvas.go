package scene

import (
	"github.com/patricktcoakley/go-rtiow/internal/geometry"
	"github.com/patricktcoakley/go-rtiow/internal/math"
	"image/color"
)

type Canvas interface {
	WritePixel(x, y int, color geometry.Vec3)
	Run()
}

type CanvasOpts struct {
	Width, Height, SamplesPerPixel int
	UsePpm                         bool
}

func NewCanvas(opts CanvasOpts) Canvas {
	if opts.UsePpm {
		return newPpmCanvas(opts.Width, opts.Height, opts.SamplesPerPixel)
	}

	return newEbitenCanvas(opts.Width, opts.Height, opts.SamplesPerPixel)
}

func newRGBAFromVec3(v geometry.Vec3, scale math.Real) color.RGBA {
	return color.RGBA{
		R: uint8(math.Clamp(0.0, 0.999, math.Sqrt(v.X*scale)) * 256),
		G: uint8(math.Clamp(0.0, 0.999, math.Sqrt(v.Y*scale)) * 256),
		B: uint8(math.Clamp(0.0, 0.999, math.Sqrt(v.Z*scale)) * 256),
		A: 255,
	}
}
