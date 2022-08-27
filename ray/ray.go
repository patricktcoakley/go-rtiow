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
