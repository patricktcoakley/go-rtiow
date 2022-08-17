package ray

import (
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
	unitDirection := r.Direction.ToUnit()
	t := 0.5 * (unitDirection[1] + 1)
	c := vec3.Vec3{1, 1, 1}.MulScalar(1 - t).Add(vec3.Vec3{0.5, 0.7, 1}.MulScalar(t))
	return c.ToRGB()
}
