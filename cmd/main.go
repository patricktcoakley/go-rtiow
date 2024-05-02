package main

import (
	"flag"
	"github.com/patricktcoakley/go-rtiow/internal/geometry"
	"github.com/patricktcoakley/go-rtiow/internal/hittable"
	"github.com/patricktcoakley/go-rtiow/internal/hittable/material"
	"github.com/patricktcoakley/go-rtiow/internal/hittable/shapes"
	"github.com/patricktcoakley/go-rtiow/internal/math"
	"github.com/patricktcoakley/go-rtiow/internal/scene"
	"github.com/patricktcoakley/go-rtiow/internal/tracer"
	"os"
	"runtime/pprof"
	"sync"
)

var (
	samplesPerPixel int
	imageWidth      int
	imageHeight     int
	aspectRatio     float64
	camera          *scene.Camera
	canvas          scene.Canvas
	canvasType      string
	world           hittable.HitList
	pgoProfile      bool
)

func samplePixel(input scene.Pixel) scene.Pixel {
	hr := new(hittable.HitRecord)

	for s := 0; s < samplesPerPixel; s++ {
		u := (math.Real(input.X) + math.Random()) / math.Real(imageWidth-1)
		v := (math.Real(input.Y) + math.Random()) / math.Real(imageHeight-1)
		r := camera.GetRay(u, v)
		input.Color = geometry.Add(input.Color, tracer.RayColor(r, world, hr))
	}

	return input
}

func randomScene() hittable.HitList {
	var mat hittable.Scatterable

	hl := make(hittable.HitList, 0, 500)
	ground := material.NewLambertian(0.5, 0.5, 0.5)
	hl = append(hl, shapes.NewSphere(0, -1000, 0, 1000, ground))

	dielectric := material.NewDielectric(1.5)
	lambertian := material.NewLambertian(0.4, 0.2, 0.1)
	metal := material.NewMetal(0.7, 0.6, 0.5, 0)

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := math.Random()
			center := geometry.Vec3{X: math.Real(a) + 0.9*math.Random(), Y: 0.2, Z: math.Real(b) + 0.9*math.Random()}
			if (center.Sub(geometry.Vec3{X: 4, Y: 0.2, Z: 0}).Length() > 0.9) {
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

				hl = append(hl, shapes.NewSphere(center.X, center.Y, center.Z, 0.2, mat))
			}
		}
	}

	hl = append(hl, shapes.NewSphere(0, 1, 0, 1, dielectric))
	hl = append(hl, shapes.NewSphere(-4, 1, 0, 1, lambertian))
	hl = append(hl, shapes.NewSphere(4, 1, 0, 1, metal))

	return hl
}

func main() {
	flag.IntVar(&samplesPerPixel, "samples", 100, "Number of samples per Pixel")
	flag.IntVar(&imageWidth, "width", 1200, "Width of render")
	flag.Float64Var(&aspectRatio, "aspect-ratio", 3.0/2.0, "The aspect ratio of render")
	flag.StringVar(&canvasType, "canvas", "png", "Canvas for the scene: 'ebiten' or 'png'")
	flag.BoolVar(&pgoProfile, "pgo-profile", false, "Enable to PGO profile")

	flag.Parse()

	if pgoProfile {
		f, err := os.Create("default.pgo")
		if err != nil {
			panic(err)
		}

		if err = pprof.StartCPUProfile(f); err != nil {
			panic(err)
		}

		defer pprof.StopCPUProfile()
	}

	world = randomScene()
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

	canvasOpts := scene.CanvasOpts{
		Width:           imageWidth,
		Height:          imageHeight,
		SamplesPerPixel: samplesPerPixel,
		CanvasType:      canvasType,
	}

	canvas = scene.NewCanvas(canvasOpts)
	pixels := make(chan scene.Pixel, imageWidth*imageHeight)

	var wg sync.WaitGroup
	wg.Add(imageWidth * imageHeight)

	go func() {
		for y := 0; y < imageHeight; y++ {
			for x := 0; x < imageWidth; x++ {
				pixels <- scene.Pixel{X: x, Y: y}
			}
		}
	}()

	go func() {
		for p := range pixels {
			go func(p scene.Pixel) {
				canvas.WritePixel(samplePixel(p))
				wg.Done()
			}(p)
		}
	}()

	wg.Wait()
	close(pixels)
	canvas.Run()
}
