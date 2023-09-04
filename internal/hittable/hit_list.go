package hittable

import (
	"github.com/patricktcoakley/go-rtiow/internal/math"
	"github.com/patricktcoakley/go-rtiow/internal/ray"
)

type HitList []Hittable

func (hl HitList) Hit(r ray.Ray, tMin, tMax math.Real, hr *HitRecord) bool {
	var tempHr HitRecord
	var hitAnything bool
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
