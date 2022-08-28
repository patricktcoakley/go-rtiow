package tracer

import (
	"math"

	"github.com/patricktcoakley/go-rtiow/internal/hittable"
	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

func RayColor(r ray.Ray, world hittable.Hittable) vec3.Vec3 {
	var hr hittable.HitRecord
	if world.Hit(r, 0, math.MaxFloat64, &hr) {
		return vec3.MulScalar(
			vec3.Add(hr.Normal, vec3.Vec3{1, 1, 1}),
			0.5,
		)
	}

	t := (r.Direction.ToUnit()[1] + 1.0) * 0.5
	return vec3.Add(
		vec3.MulScalar(vec3.Vec3{1, 1, 1}, 1.0-t),
		vec3.MulScalar(vec3.Vec3{0.5, 0.7, 1.0}, t),
	)
}
