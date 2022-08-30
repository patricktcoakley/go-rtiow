package hittable

import (
	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

type Hittable interface {
	Hit(r ray.Ray, tMin, tMax float64, hr *HitRecord) bool
}

type HitRecord struct {
	Point     vec3.Vec3
	Normal    vec3.Vec3
	Material  Material
	T         float64
	FrontFace bool
}

type HitList []Hittable

func (hr *HitRecord) SetFaceNormal(r ray.Ray, outwardNormal vec3.Vec3) {
	hr.FrontFace = vec3.Dot(r.Direction, outwardNormal) < 0
	if hr.FrontFace {
		hr.Normal = outwardNormal
	} else {
		hr.Normal = outwardNormal.Neg()
	}
}

func (hl HitList) Hit(r ray.Ray, tMin, tMax float64, hr *HitRecord) bool {
	var tempHr HitRecord
	hitAnything := false
	closestSoFar := tMax

	for _, object := range hl {
		if object.Hit(r, tMin, closestSoFar, &tempHr) {
			hitAnything = true
			closestSoFar = tempHr.T
			*hr = tempHr
		}
	}

	return hitAnything
}
