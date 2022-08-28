package camera

import (
	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

type Camera struct {
	origin          vec3.Vec3
	lowerLeftCorner vec3.Vec3
	horizontal      vec3.Vec3
	vertical        vec3.Vec3
}

func NewCamera(aspectRatio float64) Camera {
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	var origin vec3.Vec3
	horizontal := vec3.Vec3{viewportWidth, 0, 0}
	vertical := vec3.Vec3{0, viewportHeight, 0}
	lowerLeftCorner := origin.Sub(horizontal.DivScalar(2)).Sub(vertical.DivScalar(2)).Sub(vec3.Vec3{0, 0, focalLength})

	return Camera{origin, lowerLeftCorner, horizontal, vertical}
}

func (c Camera) GetRay(u, v float64) ray.Ray {
	return ray.Ray{Origin: c.origin, Direction: c.lowerLeftCorner.Add(c.horizontal.MulScalar(u)).Add(c.vertical.MulScalar(v)).Sub(c.origin)}
}
