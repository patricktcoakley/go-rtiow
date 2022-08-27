package hittable

import (
	"math"

	"github.com/patricktcoakley/go-rtiow/geom"
)

type Sphere struct {
	Center geom.Vec3
	Radius float64
}

func NewSphere(centerX, centerY, centerZ, radius float64) Sphere {
	v := geom.Vec3{centerX, centerY, centerZ}
	return Sphere{v, radius}
}

func (s Sphere) Hit(r geom.Ray, tMin, tMax float64, hr *HitRecord) bool {
	originCenter := r.Origin.Sub(s.Center)
	a := r.Direction.LengthSquared()
	half_b := geom.Dot(r.Direction, originCenter)
	c := originCenter.LengthSquared() - s.Radius*s.Radius
	discriminant := half_b*half_b - a*c
	if discriminant < 0 {
		return false
	}
	sqrtd := math.Sqrt(discriminant)
	root := (-half_b - math.Sqrt(discriminant)) / a
	if root < tMin || tMax < root {
		root = (-half_b - sqrtd) / a
		if root < tMin || tMax < root {
			return false
		}
	}
	hr.T = root
	hr.Point = r.At(hr.T)
	hr.Normal = (hr.Point.Sub(s.Center)).DivScalar(s.Radius)
	outwardNormal := (hr.Point.Sub(s.Center)).DivScalar(s.Radius)
	hr.SetFaceNormal(r, outwardNormal)
	return true
}
