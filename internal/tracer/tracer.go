package tracer

import (
	"math"

	"github.com/patricktcoakley/go-rtiow/internal/hittable"
	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

const maxDepth int = 50

func RayColor(r ray.Ray, world hittable.Hittable) vec3.Vec3 {
	var hr hittable.HitRecord
	for depth := 0; depth < maxDepth; depth++ {
		return rayColor(r, world, &hr)
	}
	return vec3.Vec3{}
}

func rayColor(r ray.Ray, world hittable.Hittable, hr *hittable.HitRecord) vec3.Vec3 {
	if world.Hit(r, 0.001, math.MaxFloat64, hr) {
		scattered := ray.Ray{}
		attenuation := vec3.Vec3{}
		if hr.Material.Scatter(r, hr, &attenuation, &scattered) {
			return vec3.Mul(rayColor(scattered, world, hr), attenuation)
		}
		return vec3.Vec3{}
	}

	t := (r.Direction.ToUnit()[1] + 1.0) * 0.5
	return vec3.Add(
		vec3.MulScalar(vec3.Vec3{1, 1, 1}, 1.0-t),
		vec3.MulScalar(vec3.Vec3{0.5, 0.7, 1.0}, t),
	)
}
