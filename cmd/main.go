package main

import (
	"flag"
	"github.com/patricktcoakley/go-rtiow/internal/camera"
	"github.com/patricktcoakley/go-rtiow/internal/canvas"
	"github.com/patricktcoakley/go-rtiow/internal/hittable"
	"github.com/patricktcoakley/go-rtiow/internal/math"
	"github.com/patricktcoakley/go-rtiow/internal/shapes"
	"github.com/patricktcoakley/go-rtiow/internal/tracer"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
	"runtime"
	"sync"
)

var samplesPerPixel int
var imageWidth int
var aspectRatio float64
var jobs int

func init() {
	flag.IntVar(&samplesPerPixel, "samples", 100, "Number of samples per pixel")
	flag.IntVar(&imageWidth, "width", 400, "Width of render")
	flag.Float64Var(&aspectRatio, "aspect-ratio", 16.0/9.0, "The aspect ratio of render")
	flag.IntVar(&jobs, "jobs", runtime.GOMAXPROCS(0), "The number of jobs")
}

func randomScene() hittable.HitList {
	world := hittable.HitList{}
	ground := hittable.NewLambertian(0.5, 0.5, 0.5)
	world = append(world, shapes.NewSphere(0, -1000, 0, 1000, ground))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := math.Random()
			center := vec3.Vec3{X: math.Real(a) + 0.9*math.Random(), Y: 0.2, Z: math.Real(b) + 0.9*math.Random()}
			if (center.Sub(vec3.Vec3{X: 4, Y: 0.2, Z: 0}).Length() > 0.9) {
				var mat hittable.Material
				if chooseMat < 0.8 {
					albedo := vec3.NewRandomVec3().Mul(vec3.NewRandomVec3())
					mat = hittable.NewLambertian(albedo.X, albedo.Y, albedo.Z)
				} else if chooseMat < 0.95 {
					albedo := vec3.NewRandomRangeVec3(0.5, 1)
					fuzz := 0.5 * math.Random()
					mat = hittable.NewMetal(albedo.X, albedo.Y, albedo.Z, fuzz)
				} else {
					mat = hittable.NewDielectric(1.5)
				}
				world = append(world, shapes.NewSphere(center.X, center.Y, center.Z, 0.2, mat))
			}
		}
	}
	mat1 := hittable.NewDielectric(1.5)
	mat2 := hittable.NewLambertian(0.4, 0.2, 0.1)
	mat3 := hittable.NewMetal(0.7, 0.6, 0.5, 0)
	world = append(world, shapes.NewSphere(0, 1, 0, 1, mat1))
	world = append(world, shapes.NewSphere(-4, 1, 0, 1, mat2))
	world = append(world, shapes.NewSphere(4, 1, 0, 1, mat3))

	return world
}

func main() {
	flag.Parse()
	aspectRatio := math.Real(aspectRatio)
	imageHeight := int(math.Real(imageWidth) / aspectRatio)
	world := randomScene()
	lookFrom := vec3.Vec3{X: 13, Y: 2, Z: 3}
	lookAt := vec3.Vec3{X: 0, Y: 0, Z: 0}
	cam := camera.NewCamera(
		lookFrom,
		lookAt,
		vec3.Vec3{X: 0, Y: 1, Z: 0},
		aspectRatio,
		20,
		0.1,
		10,
	)
	viewer := canvas.NewCanvas(imageWidth, imageHeight, samplesPerPixel)
	type coord struct {
		x, y  int
		color vec3.Vec3
	}
	coords := make(chan coord)
	pixels := make(chan coord)

	go func() {
		for y := 0; y < imageHeight; y++ {
			for x := 0; x < imageWidth; x++ {
				coords <- coord{x, y, vec3.Vec3{}}
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
					var pixelColor vec3.Vec3
					for s := 0; s < samplesPerPixel; s++ {
						u := (math.Real(c.x) + math.Random()) / math.Real(imageWidth-1)
						v := (math.Real(c.y) + math.Random()) / math.Real(imageHeight-1)
						r := cam.GetRay(u, v)
						pixelColor = vec3.Add(pixelColor, tracer.RayColor(r, world))
					}
					pixels <- coord{c.x, c.y, pixelColor}
				}
			}()
		}
		wg.Wait()
		close(pixels)
	}()

	for c := range pixels {
		viewer.WritePixel(c.x, c.y, c.color)
	}

	viewer.WriteImage()
}
