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

var samplesPerPixel int
var imageWidth int
var aspectRatio float64
var jobs int

func init() {
	flag.IntVar(&samplesPerPixel, "samples", 100, "Number of samples per pixel")
	flag.IntVar(&imageWidth, "width", 400, "Width of render")
	flag.Float64Var(&aspectRatio, "aspect-ratio", 16.0/9.0, "The aspect ratio of render")
	flag.IntVar(&jobs, "jobs", 10, "The number of jobs")
}

func randomScene() hittable.HitList {
	world := hittable.HitList{}
	ground := hittable.NewLambertian(0.5, 0.5, 0.5)
	world = append(world, shapes.NewSphere(0, -1000, 0, 1000, ground))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := vec3.Vec3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}
			if (center.Sub(vec3.Vec3{4, 0.2, 0}).Length() > 0.9) {
				var mat hittable.Material
				if chooseMat < 0.8 {
					albedo := vec3.NewRandomVec3().Mul(vec3.NewRandomVec3())
					mat = hittable.NewLambertian(albedo[0], albedo[1], albedo[2])
				} else if chooseMat < 0.95 {
					albedo := vec3.NewRandomRangeVec3(0.5, 1)
					fuzz := 0.5 * rand.Float64()
					mat = hittable.NewMetal(albedo[0], albedo[1], albedo[2], fuzz)
				} else {
					mat = hittable.NewDielectric(1.5)
				}
				world = append(world, shapes.NewSphere(center[0], center[1], center[2], 0.2, mat))
			}
		}
	}
	mats := []hittable.Material{hittable.NewDielectric(1.5), hittable.NewLambertian(rand.Float64(), rand.Float64(), rand.Float64()), hittable.NewMetal(0.7, 0.6, 0.5, 0)}
	world = append(world, shapes.NewSphere(0, 1, 0, 1, mats[rand.Intn(2)]))
	world = append(world, shapes.NewSphere(-4, 1, 0, 1, mats[rand.Intn(2)]))
	world = append(world, shapes.NewSphere(4, 1, 0, 1, mats[rand.Intn(2)]))

	return world
}

func main() {
	flag.Parse()
	imageHeight := int(float64(imageWidth) / aspectRatio)
	world := randomScene()
	lookFrom := vec3.Vec3{13, 2, 3}
	lookAt := vec3.Vec3{0, 0, 0}
	camera := camera.NewCamera(
		lookFrom,
		lookAt,
		vec3.Vec3{0, 1, 0},
		aspectRatio,
		20,
		0.1,
		10,
	)
	viewer := canvas.NewCanvas(imageWidth, imageHeight, samplesPerPixel)
	type coord struct{ x, y int }
	ch := make(chan coord)

	for job := 1; job <= jobs; job++ {
		go func() {
			for c := range ch {
				var pixelColor vec3.Vec3
				for s := 0; s < samplesPerPixel; s++ {
					u := (float64(c.x) + rand.Float64()) / float64(imageWidth-1)
					v := (float64(c.y) + rand.Float64()) / float64(imageHeight-1)
					r := camera.GetRay(u, v)
					pixelColor = vec3.Add(pixelColor, tracer.RayColor(r, world))
				}
				viewer.WritePixel(c.x, c.y, pixelColor)
			}
		}()
	}

	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			ch <- coord{x, y}
		}
	}

	viewer.WriteImage()
}