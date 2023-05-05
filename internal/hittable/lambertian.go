package hittable

import (
	"github.com/patricktcoakley/go-rtiow/internal/math"
	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

type Lambertian struct {
	Albedo vec3.Vec3
}

func NewLambertian(r, g, b math.Real) *Lambertian {
	return &Lambertian{vec3.Vec3{X: r, Y: g, Z: b}}
}

func (l *Lambertian) Scatter(rIn ray.Ray, hr *HitRecord, attenuation *vec3.Vec3, scattered *ray.Ray) bool {
	scatterDirection := hr.Normal.Add(vec3.NewRandomUnitVector())
	if scatterDirection.NearZero() {
		scatterDirection = hr.Normal
	}

	*scattered = ray.Ray{Origin: hr.Point, Direction: scatterDirection}
	*attenuation = l.Albedo

	return true
}
