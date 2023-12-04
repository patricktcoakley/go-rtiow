package main

import (
	"flag"
	"github.com/patricktcoakley/go-rtiow/internal/geometry"
	"github.com/patricktcoakley/go-rtiow/internal/hittable"
	"github.com/patricktcoakley/go-rtiow/internal/hittable/material"
	"github.com/patricktcoakley/go-rtiow/internal/math"
	"github.com/patricktcoakley/go-rtiow/internal/scene"
	"github.com/patricktcoakley/go-rtiow/internal/shapes"
	"github.com/patricktcoakley/go-rtiow/internal/tracer"
	"sync"
)

var samplesPerPixel int
var imageWidth int
var imageHeight int
var aspectRatio float64
var camera scene.Camera
var canvas scene.Canvas
var world hittable.HitList

func init() {
	flag.IntVar(&samplesPerPixel, "samples", 1000, "Number of samples per pixel")
	flag.IntVar(&imageWidth, "width", 1200, "Width of render")
	flag.Float64Var(&aspectRatio, "aspect-ratio", 3.0/2.0, "The aspect ratio of render")
}

func samplePixel(x int, y int) pixel {
	var pixelColor geometry.Vec3
	for s := 0; s < samplesPerPixel; s++ {
		u := (math.Real(x) + math.Random()) / math.Real(imageWidth-1)
		v := (math.Real(y) + math.Random()) / math.Real(imageHeight-1)
		r := camera.GetRay(u, v)
		pixelColor = geometry.Add(pixelColor, tracer.RayColor(r, world))
	}

	return pixel{x, y, pixelColor}
}

func randomScene() hittable.HitList {
	world := hittable.HitList{}
	ground := material.NewLambertian(0.5, 0.5, 0.5)
	world = append(world, shapes.NewSphere(0, -1000, 0, 1000, ground))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := math.Random()
			center := geometry.Vec3{X: math.Real(a) + 0.9*math.Random(), Y: 0.2, Z: math.Real(b) + 0.9*math.Random()}
			if (center.Sub(geometry.Vec3{X: 4, Y: 0.2, Z: 0}).Length() > 0.9) {
				var mat hittable.Scatterable
				if chooseMat < 0.8 {
					albedo := geometry.NewRandomVec3().Mul(geometry.NewRandomVec3())
					mat = material.NewLambertian(albedo.X, albedo.Y, albedo.Z)
				} else if chooseMat < 0.95 {
					albedo := geometry.NewRandomRangeVec3(0.5, 1)
					fuzz := 0.5 * math.Random()
					mat = material.NewMetal(albedo.X, albedo.Y, albedo.Z, fuzz)
				} else {
					mat = material.NewDielectric(1.5)
				}
				world = append(world, shapes.NewSphere(center.X, center.Y, center.Z, 0.2, mat))
			}
		}
	}

	dielectric := material.NewDielectric(1.5)
	lambertian := material.NewLambertian(0.4, 0.2, 0.1)
	metal := material.NewMetal(0.7, 0.6, 0.5, 0)
	world = append(world, shapes.NewSphere(0, 1, 0, 1, dielectric))
	world = append(world, shapes.NewSphere(-4, 1, 0, 1, lambertian))
	world = append(world, shapes.NewSphere(4, 1, 0, 1, metal))

	return world
}

type pixel struct {
	x, y  int
	color geometry.Vec3
}

func main() {
	flag.Parse()
	aspectRatio := math.Real(aspectRatio)
	imageHeight = int(math.Real(imageWidth) / aspectRatio)
	lookFrom := geometry.Vec3{X: 13, Y: 2, Z: 3}
	lookAt := geometry.Vec3{}
	verticalUp := geometry.Vec3{Y: 1}
	verticalFov := math.Real(20)
	aperture := math.Real(0.1)
	focusDistance := math.Real(10.0)
	camera = scene.NewCamera(
		lookFrom,
		lookAt,
		verticalUp,
		aspectRatio,
		verticalFov,
		aperture,
		focusDistance,
	)
	world = randomScene()
	canvas = scene.NewCanvas(imageWidth, imageHeight, samplesPerPixel)
	pixels := make(chan pixel, imageWidth*imageHeight)

	var wg sync.WaitGroup
	wg.Add(imageWidth * imageHeight)

	go func() {
		for p := range pixels {
			canvas.WritePixel(p.x, p.y, p.color)
			wg.Done()
		}
	}()

	for x := 0; x < imageWidth; x++ {
		go func(x int) {
			for y := 0; y < imageHeight; y++ {
				pixels <- samplePixel(x, y)
			}
		}(x)
	}

	wg.Wait()
	close(pixels)

	canvas.WriteImage()
}
