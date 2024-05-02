package scene

import (
	"fmt"
	"github.com/patricktcoakley/go-rtiow/internal/geometry"
	"github.com/patricktcoakley/go-rtiow/internal/math"
	"image/color"
)

type Canvas interface {
	WritePixel(pixel Pixel)
	Run()
}

type CanvasOpts struct {
	Width, Height, SamplesPerPixel int
	CanvasType                     string
}

func NewCanvas(opts CanvasOpts) Canvas {
	switch opts.CanvasType {
	case "png":
		return newPngCanvas(opts.Width, opts.Height, opts.SamplesPerPixel)
	case "ebiten":
		return newEbitenCanvas(opts.Width, opts.Height, opts.SamplesPerPixel)
	default:
		panic(fmt.Sprintf("invalid canvas type: got %s, expected ['ebiten' | 'png']", opts.CanvasType))
	}
}

func newRGBAFromVec3(v geometry.Vec3, scale math.Real) color.RGBA {
	rr := math.Sqrt(v.X * scale)
	gg := math.Sqrt(v.Y * scale)
	bb := math.Sqrt(v.Z * scale)

	r := math.Clamp(0, 255, uint8(rr*255))
	g := math.Clamp(0, 255, uint8(gg*255))
	b := math.Clamp(0, 255, uint8(bb*255))

	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}
