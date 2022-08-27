package main

import (
	"math"
	"math/rand"

	"github.com/patricktcoakley/go-rtiow/internal/camera"
	"github.com/patricktcoakley/go-rtiow/internal/canvas"
	"github.com/patricktcoakley/go-rtiow/internal/hittable"
	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

func rayColor(r ray.Ray, world hittable.Hittable) vec3.Vec3 {
	var hr hittable.HitRecord
	if world.Hit(r, 0, math.MaxFloat64, &hr) {
		return vec3.MulScalar(
			vec3.Add(hr.Normal, vec3.Vec3{1, 1, 1}),
			0.5,
		)
	}

	t := (r.Direction.ToUnit()[1] + 1.0) * 0.5
	return vec3.Add(
		vec3.MulScalar(vec3.Vec3{1, 1, 1}, 1.0-t),
		vec3.MulScalar(vec3.Vec3{0.5, 0.7, 1.0}, t),
	)
}

func main() {
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)
	world := hittable.HittableList{hittable.NewSphere(0, 0, -1, 0.5), hittable.NewSphere(0, -100.5, -1, 100)}
	camera := camera.NewCamera()
	viewer := canvas.NewCanvas(imageWidth, imageHeight, "go-rtiow")

	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			u := (float64(x) + rand.Float64()) / float64(imageWidth-1)
			v := (float64(y) + rand.Float64()) / float64(imageHeight-1)
			r := camera.GetRay(u, v)
			color := rayColor(r, world)
			viewer.WritePixel(x, y, color)
		}
	}
	viewer.Run()
}
