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
	"runtime"
	"sync"
)

var samplesPerPixel int
var imageWidth int
var aspectRatio float64
var jobs int

func init() {
	flag.IntVar(&samplesPerPixel, "samples", 1000, "Number of samples per pixel")
	flag.IntVar(&imageWidth, "width", 1200, "Width of render")
	flag.Float64Var(&aspectRatio, "aspect-ratio", 3.0/2.0, "The aspect ratio of render")
	flag.IntVar(&jobs, "jobs", runtime.GOMAXPROCS(0), "The number of jobs")
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

type coord struct{ x, y int }

type pixel struct {
	coord
	color geometry.Vec3
}

func main() {
	flag.Parse()
	aspectRatio := math.Real(aspectRatio)
	imageHeight := int(math.Real(imageWidth) / aspectRatio)
	world := randomScene()
	lookFrom := geometry.Vec3{X: 13, Y: 2, Z: 3}
	lookAt := geometry.Vec3{}
	cam := scene.NewCamera(
		lookFrom,
		lookAt,
		geometry.Vec3{Y: 1},
		aspectRatio,
		20,
		0.1,
		10,
	)
	viewer := scene.NewCanvas(imageWidth, imageHeight, samplesPerPixel)

	coords := make(chan coord)
	pixels := make(chan pixel)

	go func() {
		for y := 0; y < imageHeight; y++ {
			for x := 0; x < imageWidth; x++ {
				coords <- coord{x, y}
			}
		}
		close(coords)
	}()

	go func() {
		var wg sync.WaitGroup
		wg.Add(jobs)
		for job := 1; job <= jobs; job++ {
			go func() {
				defer wg.Done()
				for c := range coords {
					var pixelColor geometry.Vec3
					for s := 0; s < samplesPerPixel; s++ {
						u := (math.Real(c.x) + math.Random()) / math.Real(imageWidth-1)
						v := (math.Real(c.y) + math.Random()) / math.Real(imageHeight-1)
						r := cam.GetRay(u, v)
						pixelColor = geometry.Add(pixelColor, tracer.RayColor(r, world))
					}
					pixels <- pixel{coord{c.x, c.y}, pixelColor}
				}
			}()
		}
		wg.Wait()
		close(pixels)
	}()

	for p := range pixels {
		viewer.WritePixel(p.x, p.y, p.color)
	}

	viewer.WriteImage()
}
