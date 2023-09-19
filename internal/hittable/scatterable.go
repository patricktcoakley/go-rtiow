package hittable

import (
	"github.com/patricktcoakley/go-rtiow/internal/geometry"
)

type Scatterable interface {
	Scatter(rIn geometry.Ray, hr *HitRecord, attenuation *geometry.Vec3, scattered *geometry.Ray) bool
}
