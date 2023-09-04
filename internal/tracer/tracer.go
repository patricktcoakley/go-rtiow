package tracer

import (
	"github.com/patricktcoakley/go-rtiow/internal/hittable"
	"github.com/patricktcoakley/go-rtiow/internal/math"
	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

const maxDepth int = 50

func RayColor(r ray.Ray, obj hittable.Hittable) vec3.Vec3 {
	return rayColor(r, obj, &hittable.HitRecord{}, maxDepth)
}

func rayColor(r ray.Ray, obj hittable.Hittable, hr *hittable.HitRecord, depth int) vec3.Vec3 {
	if depth <= 0 {
		return vec3.Vec3{}
	}
	if obj.Hit(r, math.MinReal, math.MaxReal, hr) {
		var scattered ray.Ray
		var attenuation vec3.Vec3
		if hr.Material.Scatter(r, hr, &attenuation, &scattered) {
			return vec3.Mul(attenuation, rayColor(scattered, obj, hr, depth-1))
		}
		return vec3.Vec3{}
	}

	t := (r.Direction.ToUnit().Y + 1.0) * 0.5
	return vec3.Add(
		vec3.MulScalar(vec3.Vec3{X: 1, Y: 1, Z: 1}, 1.0-t),
		vec3.MulScalar(vec3.Vec3{X: 0.5, Y: 0.7, Z: 1.0}, t),
	)
}
