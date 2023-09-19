package geometry

import (
	"github.com/patricktcoakley/go-rtiow/internal/math"
)

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (r Ray) At(t math.Real) Vec3 {
	return Add(r.Origin, MulScalar(r.Direction, t))
}
