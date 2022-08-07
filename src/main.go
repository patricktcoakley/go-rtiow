package main

import (
	"fmt"
	"os"
)

const (
	Width  = 256
	Height = 256
)

func main() {
	f, err := os.Create("out.ppm")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", Width, Height))

	for j := Height; j >= 0; j-- {
		fmt.Printf("\rScanlines remaining: %d\n", j)
		for i := 0; i < Width; i++ {
			r := float32(i) / (Width - 1)
			g := float32(j) / (Height - 1)
			b := 0.25
			f.WriteString(fmt.Sprintf("%d %d %d\n", int(r*255.99), int(g*255.99), int(b*255.99)))
		}
	}
}
