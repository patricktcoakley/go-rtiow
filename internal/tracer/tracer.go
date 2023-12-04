package tracer

import (
	"github.com/patricktcoakley/go-rtiow/internal/geometry"
	"github.com/patricktcoakley/go-rtiow/internal/hittable"
	"github.com/patricktcoakley/go-rtiow/internal/math"
)

const maxDepth int = 50

var (
	white = geometry.Vec3{X: 1, Y: 1, Z: 1}
	blue  = geometry.Vec3{X: 0.5, Y: 0.7, Z: 1.0}
)

func RayColor(r geometry.Ray, obj hittable.Hittable) geometry.Vec3 {
	return rayColor(r, obj, &hittable.HitRecord{}, maxDepth)
}

func rayColor(r geometry.Ray, obj hittable.Hittable, hr *hittable.HitRecord, depth int) geometry.Vec3 {
	if depth <= 0 {
		return geometry.Vec3{}
	}

	if !obj.Hit(r, math.MinReal, math.MaxReal, hr) {
		return geometry.Add(
			geometry.MulScalar(white, 1.0-((r.Direction.ToUnit().Y+1.0)*0.5)),
			geometry.MulScalar(blue, (r.Direction.ToUnit().Y+1.0)*0.5),
		)
	}

	return handleHit(r, obj, hr, depth)
}

func handleHit(r geometry.Ray, obj hittable.Hittable, hr *hittable.HitRecord, depth int) geometry.Vec3 {
	var scattered geometry.Ray
	var attenuation geometry.Vec3
	if hr.Material.Scatter(r, hr, &attenuation, &scattered) {
		return geometry.Mul(attenuation, rayColor(scattered, obj, hr, depth-1))
	}
	return geometry.Vec3{}
}
