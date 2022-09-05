package tracer

import (
	"math"

	"github.com/patricktcoakley/go-rtiow/internal/hittable"
	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

const maxDepth int = 50

func RayColor(r ray.Ray, obj hittable.Hittable) vec3.Vec3 {
	var hr hittable.HitRecord
	return rayColor(r, obj, &hr, maxDepth-1)
}

func rayColor(r ray.Ray, obj hittable.Hittable, hr *hittable.HitRecord, depth int) vec3.Vec3 {
	if depth <= 0 {
		return vec3.Vec3{}
	}
	if obj.Hit(r, 0.001, math.MaxFloat64, hr) {
		scattered := ray.Ray{}
		attenuation := vec3.Vec3{}
		if hr.Material.Scatter(r, hr, &attenuation, &scattered) {
			return vec3.Mul(attenuation, rayColor(scattered, obj, hr, depth-1))
		}
		return vec3.Vec3{}
	}

	t := (r.Direction.ToUnit()[1] + 1.0) * 0.5
	return vec3.Add(
		vec3.MulScalar(vec3.Vec3{1, 1, 1}, 1.0-t),
		vec3.MulScalar(vec3.Vec3{0.5, 0.7, 1.0}, t),
	)
}
