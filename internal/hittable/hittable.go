package hittable

import (
	"github.com/patricktcoakley/go-rtiow/internal/ray"
)

type Hittable interface {
	Hit(r ray.Ray, tMin, tMax float64, hr *HitRecord) bool
}
