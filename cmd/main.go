package main

import (
	"github.com/patricktcoakley/go-rtiow/color"
	"github.com/patricktcoakley/go-rtiow/game"
)

const (
	Width  = 256
	Height = 256
)

func main() {
	game := game.New(Width, Height, "go-rtiow")
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			r := uint8(float64(x) / (Width - 1) * 255.99)
			g := uint8(float64(y) / (Height - 1) * 255.99)
			b := uint8(64)
			rgb := color.New(r, g, b)
			game.WritePixel(x, y, rgb)
		}
	}
	game.Run()
}
