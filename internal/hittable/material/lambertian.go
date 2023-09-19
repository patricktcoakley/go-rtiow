package material

import (
	"github.com/patricktcoakley/go-rtiow/internal/geometry"
	"github.com/patricktcoakley/go-rtiow/internal/hittable"
	"github.com/patricktcoakley/go-rtiow/internal/math"
)

type Lambertian struct {
	Albedo geometry.Vec3
}

func NewLambertian(r, g, b math.Real) *Lambertian {
	return &Lambertian{geometry.Vec3{X: r, Y: g, Z: b}}
}

func (l *Lambertian) Scatter(_ geometry.Ray, hr *hittable.HitRecord, attenuation *geometry.Vec3, scattered *geometry.Ray) bool {
	scatterDirection := hr.Normal.Add(geometry.NewRandomUnitVector())
	if scatterDirection.NearZero() {
		scatterDirection = hr.Normal
	}

	*scattered = geometry.Ray{Origin: hr.Point, Direction: scatterDirection}
	*attenuation = l.Albedo

	return true
}
