package hittable

import (
	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

type Lambertian struct {
	Albedo vec3.Vec3
}

func NewLambertian(r, g, b float64) *Lambertian {
	return &Lambertian{vec3.Vec3{X: r, Y: g, Z: b}}
}

func (l *Lambertian) Scatter(rIn ray.Ray, hr *HitRecord, attenuation *vec3.Vec3, scattered *ray.Ray) bool {
	scatter_direction := hr.Normal.Add(vec3.NewRandomUnitVector())
	if scatter_direction.NearZero() {
		scatter_direction = hr.Normal
	}

	*scattered = ray.Ray{Origin: hr.Point, Direction: scatter_direction}
	*attenuation = l.Albedo

	return true
}
