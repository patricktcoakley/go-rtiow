package ray

import (
	"math"

	"github.com/patricktcoakley/go-rtiow/vec3"
)

type Ray struct {
	Origin    vec3.Vec3
	Direction vec3.Vec3
}

func (r Ray) At(t float64) vec3.Vec3 {
	return vec3.Add(r.Origin, vec3.MulScalar(r.Direction, t))
}

func (r Ray) Color() vec3.Color {
	t := hitSphere(vec3.Vec3{0, 0, -1}, 0.5, r)
	if t > 0.0 {
		n := r.At(t).Sub(vec3.Vec3{0, 0, -1}).ToUnit()
		return vec3.Vec3{n[0] + 1, n[1] + 1, n[2] + 1}.MulScalar(0.5).ToRGB()
	}
	unitDirection := r.Direction.ToUnit()
	t = 0.5 * (unitDirection[1] + 1)
	c := vec3.Vec3{1, 1, 1}.MulScalar(1 - t).Add(vec3.Vec3{0.5, 0.7, 1}.MulScalar(t))
	return c.ToRGB()
}

func hitSphere(center vec3.Vec3, radius float64, r Ray) float64 {
	originCenter := r.Origin.Sub(center)
	a := vec3.Dot(r.Direction, r.Direction)
	b := 2.0 * vec3.Dot(originCenter, r.Direction)
	c := vec3.Dot(originCenter, originCenter) - (radius * radius)
	discriminant := b*b - (4 * a * c)
	if discriminant < 0 {
		return -1.0
	}
	return (-b - math.Sqrt(discriminant)) / (2.0 * a)
}
