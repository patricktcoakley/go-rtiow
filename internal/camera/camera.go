package camera

import (
	"math"

	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

type Camera struct {
	origin          vec3.Vec3
	lowerLeftCorner vec3.Vec3
	horizontal      vec3.Vec3
	vertical        vec3.Vec3
}

func NewCamera(
	lookFrom,
	lookAt,
	verticalUp vec3.Vec3,
	aspectRatio,
	verticalFov float64) *Camera {

	theta := degreesToRadians(verticalFov)
	h := math.Tan(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight
	w := lookFrom.Sub(lookAt).ToUnit()
	u := vec3.Cross(verticalUp, w).ToUnit()
	v := vec3.Cross(w, u)

	origin := lookFrom
	horizontal := u.MulScalar(viewportWidth)
	vertical := v.MulScalar(viewportHeight)
	lowerLeftCorner := origin.Sub(horizontal.DivScalar(2)).Sub(vertical.DivScalar(2)).Sub(w)

	return &Camera{origin, lowerLeftCorner, horizontal, vertical}
}

func (c *Camera) GetRay(s, t float64) ray.Ray {
	return ray.Ray{
		Origin:    c.origin,
		Direction: c.lowerLeftCorner.Add(c.horizontal.MulScalar(s).Add(c.vertical.MulScalar(t)).Sub(c.origin)),
	}
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}
