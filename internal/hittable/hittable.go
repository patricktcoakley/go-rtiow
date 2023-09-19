package hittable

import (
	"github.com/patricktcoakley/go-rtiow/internal/geometry"
	"github.com/patricktcoakley/go-rtiow/internal/math"
)

type Hittable interface {
	Hit(r geometry.Ray, tMin, tMax math.Real, hr *HitRecord) bool
}
