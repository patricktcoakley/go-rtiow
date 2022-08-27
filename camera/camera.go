package camera

import (
	"github.com/patricktcoakley/go-rtiow/geom"
)

type Camera struct {
	origin          geom.Vec3
	lowerLeftCorner geom.Vec3
	horizontal      geom.Vec3
	vertical        geom.Vec3
}

func NewCamera() Camera {
	aspectRatio := 16.0 / 9.0
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	var origin geom.Vec3
	horizontal := geom.Vec3{viewportWidth, 0, 0}
	vertical := geom.Vec3{0, viewportHeight, 0}
	lowerLeftCorner := origin.Sub(horizontal.DivScalar(2)).Sub(vertical.DivScalar(2)).Sub(geom.Vec3{0, 0, focalLength})

	return Camera{origin, lowerLeftCorner, horizontal, vertical}
}

func (c Camera) GetRay(u, v float64) geom.Ray {
	return geom.Ray{Origin: c.origin, Direction: c.lowerLeftCorner.Add(c.horizontal.MulScalar(u)).Add(c.vertical.MulScalar(v)).Sub(c.origin)}
}
