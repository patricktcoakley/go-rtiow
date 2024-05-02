package shapes

import (
	"github.com/patricktcoakley/go-rtiow/internal/geometry"
	"github.com/patricktcoakley/go-rtiow/internal/hittable"
	"github.com/patricktcoakley/go-rtiow/internal/math"
)

type Sphere struct {
	Center   geometry.Vec3
	Radius   math.Real
	Material hittable.Scatterable
}

func NewSphere(centerX, centerY, centerZ, radius math.Real, material hittable.Scatterable) *Sphere {
	v := geometry.Vec3{X: centerX, Y: centerY, Z: centerZ}
	return &Sphere{v, radius, material}
}

func (s *Sphere) Hit(r geometry.Ray, tMin, tMax math.Real, hr *hittable.HitRecord) bool {
	originCenter := r.Origin.Sub(s.Center)
	a := r.Direction.LengthSquared()
	halfB := geometry.Dot(r.Direction, originCenter)
	c := originCenter.LengthSquared() - s.Radius*s.Radius
	discriminant := halfB*halfB - a*c

	if discriminant < 0 {
		return false
	}

	sqrtD := math.Sqrt(discriminant)
	root := (-halfB - sqrtD) / a

	if root < tMin || root > tMax {
		root = (-halfB - sqrtD) / a
		if root < tMin || root > tMax {
			return false
		}
	}

	hr.T = root
	hr.Point = r.At(hr.T)
	outwardNormal := (hr.Point.Sub(s.Center)).DivScalar(s.Radius)
	hr.SetFaceNormal(r, outwardNormal)
	hr.Material = s.Material

	return true
}
