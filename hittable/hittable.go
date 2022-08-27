package hittable

import (
	"github.com/patricktcoakley/go-rtiow/geom"
)

type Hittable interface {
	Hit(r geom.Ray, tMin, tMax float64, hr *HitRecord) bool
}

type HitRecord struct {
	Point     geom.Vec3
	Normal    geom.Vec3
	T         float64
	FrontFace bool
}

type HittableList []Hittable

func (hr *HitRecord) SetFaceNormal(r geom.Ray, outwardNormal geom.Vec3) {
	hr.FrontFace = geom.Dot(r.Direction, outwardNormal) < 0
	if hr.FrontFace {
		hr.Normal = outwardNormal
	} else {
		hr.Normal = outwardNormal
	}
}

func (hl HittableList) Hit(r geom.Ray, tMin, tMax float64, hr *HitRecord) bool {
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
