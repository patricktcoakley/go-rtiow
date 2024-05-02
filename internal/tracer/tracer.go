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

func RayColor(r geometry.Ray, obj hittable.Hittable, hr *hittable.HitRecord) geometry.Vec3 {
	var scattered geometry.Ray

	color := white
	attenuation := white
	currentRay := r

	for range maxDepth {
		didHitOrScatter := obj.Hit(currentRay, 0.001, math.MaxReal, hr) && hr.Material.Scatter(currentRay, hr, &attenuation, &scattered)

		if !(didHitOrScatter) {
			unitDirection := currentRay.Direction.ToUnit()
			a := 0.5 * (unitDirection.Y + 1.0)
			return color.Mul(white.MulScalar(1.0 - a).Add(blue.MulScalar(a)))
		}

		color = color.Mul(attenuation)
		currentRay = scattered
	}

	return geometry.Vec3{}
}
