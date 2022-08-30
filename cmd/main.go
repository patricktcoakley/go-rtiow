package main

import (
	"flag"
	"math/rand"

	"github.com/patricktcoakley/go-rtiow/internal/camera"
	"github.com/patricktcoakley/go-rtiow/internal/canvas"
	"github.com/patricktcoakley/go-rtiow/internal/hittable"
	"github.com/patricktcoakley/go-rtiow/internal/shapes"
	"github.com/patricktcoakley/go-rtiow/internal/tracer"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

const maxDepth = 50

var samplesPerPixel int
var imageWidth int
var aspectRatio float64

func init() {
	flag.IntVar(&samplesPerPixel, "samples", 100, "Number of samples per pixel")
	flag.IntVar(&imageWidth, "width", 400, "Width of render")
	flag.Float64Var(&aspectRatio, "aspect-ratio", 16.0/9.0, "The aspect ratio of render")
}

func main() {
	flag.Parse()
	imageHeight := int(float64(imageWidth) / aspectRatio)
	ground := hittable.NewLambertian(0.8, 0.8, 0.0)
	center := hittable.NewLambertian(0.7, 0.3, 0.3)
	left := hittable.NewMetal(0.8, 0.8, 0.8, 0.3)
	right := hittable.NewMetal(0.8, 0.6, 0.2, 1.0)

	world := hittable.HitList{
		shapes.NewSphere(0, -100.5, -1, 100, ground),
		shapes.NewSphere(0, 0, -1, 0.5, center),
		shapes.NewSphere(-1., 0, -1, 0.5, left),
		shapes.NewSphere(1., 0, -1, 0.5, right),
	}
	camera := camera.NewCamera(aspectRatio)
	viewer := canvas.NewCanvas(imageWidth, imageHeight, samplesPerPixel, "go-rtiow")

	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			var pixelColor vec3.Vec3
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(x) + rand.Float64()) / float64(imageWidth-1)
				v := (float64(y) + rand.Float64()) / float64(imageHeight-1)
				r := camera.GetRay(u, v)
				pixelColor = vec3.Add(pixelColor, tracer.RayColor(r, world, maxDepth))
			}
			viewer.WritePixel(x, y, pixelColor)
		}
	}
	viewer.Run()
}
