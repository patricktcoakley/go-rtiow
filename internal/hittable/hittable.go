package hittable

import (
	"github.com/patricktcoakley/go-rtiow/internal/math"
	"github.com/patricktcoakley/go-rtiow/internal/ray"
)

type Hittable interface {
	Hit(r ray.Ray, tMin, tMax math.Real, hr *HitRecord) bool
}
