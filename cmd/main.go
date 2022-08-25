package main

import (
	"github.com/patricktcoakley/go-rtiow/game"
	"github.com/patricktcoakley/go-rtiow/ray"
	"github.com/patricktcoakley/go-rtiow/vec3"
)


func main() {
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)

	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := vec3.Vec3{}
	horizontal := vec3.Vec3{viewportWidth}
	vertical := vec3.Vec3{0, viewportHeight, 0}
	lowerLeftCorner := origin.Sub(horizontal.DivScalar(2)).Sub(vertical.DivScalar(2)).Sub(vec3.Vec3{0, 0, focalLength})

	game := game.New(imageWidth, imageHeight, "go-rtiow")
	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			u := float64(x) / float64(imageWidth-1)
			v := float64(y) / float64(imageHeight-1)
			r := ray.Ray{Origin: origin, Direction: lowerLeftCorner.Add(horizontal.MulScalar(u)).Add(vertical.MulScalar(v)).Sub(origin)}

			game.WritePixel(x, y, r.Color())
		}
	}
	game.Run()
}
