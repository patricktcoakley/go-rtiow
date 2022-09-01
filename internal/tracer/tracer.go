package tracer

import (
	"math"

	"github.com/patricktcoakley/go-rtiow/internal/hittable"
	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

func RayColor(r ray.Ray, world hittable.Hittable, depth int) vec3.Vec3 {
	var hr hittable.HitRecord
	if world.Hit(r, math.SmallestNonzeroFloat64, math.MaxFloat64, &hr) {
		scattered := ray.Ray{}
		attenuation := vec3.Vec3{}
		if hr.Material.Scatter(r, &hr, &attenuation, &scattered) {
			return attenuation.Mul(RayColor(scattered, world, depth-1))
		}

		return vec3.Vec3{}
	}

	t := (r.Direction.ToUnit()[1] + 1.0) * 0.5
	return vec3.Add(
		vec3.MulScalar(vec3.Vec3{1, 1, 1}, 1.0-t),
		vec3.MulScalar(vec3.Vec3{0.5, 0.7, 1.0}, t),
	)
}
