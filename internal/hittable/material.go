package hittable

import (
	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

type Material interface {
	Scatter(rIn ray.Ray, hr *HitRecord, attenuation *vec3.Vec3, scattered *ray.Ray) bool
}
