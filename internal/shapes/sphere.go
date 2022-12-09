package shapes

import (
	"math"

	"github.com/patricktcoakley/go-rtiow/internal/hittable"
	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

type Sphere struct {
	Center   vec3.Vec3
	Radius   float64
	Material hittable.Material
}

func NewSphere(centerX, centerY, centerZ, radius float64, material hittable.Material) *Sphere {
	v := vec3.Vec3{X: centerX, Y: centerY, Z: centerZ}
	return &Sphere{v, radius, material}
}

func (s *Sphere) Hit(r ray.Ray, tMin, tMax float64, hr *hittable.HitRecord) bool {
	originCenter := r.Origin.Sub(s.Center)
	a := r.Direction.LengthSquared()
	half_b := vec3.Dot(r.Direction, originCenter)
	c := originCenter.LengthSquared() - s.Radius*s.Radius
	discriminant := half_b*half_b - a*c
	if discriminant < 0 {
		return false
	}
	sqrtd := math.Sqrt(discriminant)
	root := (-half_b - sqrtd) / a
	if root < tMin || root > tMax {
		root = (-half_b - sqrtd) / a
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
