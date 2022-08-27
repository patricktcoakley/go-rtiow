package main

import (
	"math"
	"math/rand"

	"github.com/patricktcoakley/go-rtiow/camera"
	"github.com/patricktcoakley/go-rtiow/game"
	"github.com/patricktcoakley/go-rtiow/hittable"
	"github.com/patricktcoakley/go-rtiow/ray"
	"github.com/patricktcoakley/go-rtiow/vec3"
)

func rayColor(r ray.Ray, world hittable.Hittable) vec3.Color {
	var hr hittable.HitRecord
	if world.Hit(r, 0, math.MaxFloat64, &hr) {
		return (hr.Normal.Add(vec3.Vec3{1, 1, 1})).MulScalar(0.5).ToRGB()
	}

	t := (r.Direction.ToUnit()[1] + 1.0) * 0.5
	c := (vec3.Vec3{1, 1, 1}.MulScalar(1 - t)).Add((vec3.Vec3{0.5, 0.7, 1.0}).MulScalar(t))
	return c.ToRGB()
}

func main() {
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)
	world := hittable.HittableList{hittable.NewSphere(0, 0, -1, 0.5), hittable.NewSphere(0, -100.5, -1, 100)}
	camera := camera.NewCamera()
	game := game.New(imageWidth, imageHeight, "go-rtiow")
	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			u := (float64(x) + rand.Float64()) / float64(imageWidth-1)
			v := (float64(y) + rand.Float64()) / float64(imageHeight-1)
			r := camera.GetRay(u, v)
			color := rayColor(r, world)
			game.WritePixel(x, y, color)
		}
	}
	game.Run()
}
