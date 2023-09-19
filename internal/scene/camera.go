package scene

import (
	"github.com/patricktcoakley/go-rtiow/internal/geometry"
	"github.com/patricktcoakley/go-rtiow/internal/math"
)

type Camera struct {
	origin          geometry.Vec3
	lowerLeftCorner geometry.Vec3
	horizontal      geometry.Vec3
	vertical        geometry.Vec3
	u               geometry.Vec3
	v               geometry.Vec3
	w               geometry.Vec3
	lensRadius      math.Real
}

func NewCamera(
	lookFrom,
	lookAt,
	verticalUp geometry.Vec3,
	aspectRatio,
	verticalFov,
	aperture,
	focusDistance math.Real) *Camera {

	theta := degreesToRadians(verticalFov)
	h := math.Tan(theta / 2.0)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := lookFrom.Sub(lookAt).ToUnit()
	u := geometry.Cross(verticalUp, w).ToUnit()
	v := geometry.Cross(w, u)

	origin := lookFrom
	horizontal := u.MulScalar(viewportWidth).MulScalar(focusDistance)
	vertical := v.MulScalar(viewportHeight).MulScalar(focusDistance)
	lowerLeftCorner := origin.Sub(horizontal.DivScalar(2)).Sub(vertical.DivScalar(2)).Sub(w.MulScalar(focusDistance))
	lensRadius := aperture / 2
	return &Camera{
		origin,
		lowerLeftCorner,
		horizontal,
		vertical,
		u,
		v,
		w,
		lensRadius,
	}
}

func (c *Camera) GetRay(s, t math.Real) geometry.Ray {
	rd := geometry.RandomInUnitDisk().MulScalar(c.lensRadius)
	offset := c.u.MulScalar(rd.X).Add(c.v.MulScalar(rd.Y))
	return geometry.Ray{
		Origin:    c.origin.Add(offset),
		Direction: c.lowerLeftCorner.Add(c.horizontal.MulScalar(s).Add(c.vertical.MulScalar(t)).Sub(c.origin)).Sub(offset),
	}
}

func degreesToRadians(degrees math.Real) math.Real {
	return degrees * math.Pi / 180.0
}
