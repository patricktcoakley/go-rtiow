package camera

import (
	"github.com/patricktcoakley/go-rtiow/internal/math"
	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

type Camera struct {
	origin          vec3.Vec3
	lowerLeftCorner vec3.Vec3
	horizontal      vec3.Vec3
	vertical        vec3.Vec3
	u               vec3.Vec3
	v               vec3.Vec3
	w               vec3.Vec3
	lensRadius      math.Real
}

func NewCamera(
	lookFrom,
	lookAt,
	verticalUp vec3.Vec3,
	aspectRatio,
	verticalFov,
	aperture,
	focusDistance math.Real) *Camera {

	theta := degreesToRadians(verticalFov)
	h := math.Tan(theta / 2.0)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := lookFrom.Sub(lookAt).ToUnit()
	u := vec3.Cross(verticalUp, w).ToUnit()
	v := vec3.Cross(w, u)

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

func (c *Camera) GetRay(s, t math.Real) ray.Ray {
	rd := vec3.RandomInUnitDisk().MulScalar(c.lensRadius)
	offset := c.u.MulScalar(rd.X).Add(c.v.MulScalar(rd.Y))
	return ray.Ray{
		Origin:    c.origin.Add(offset),
		Direction: c.lowerLeftCorner.Add(c.horizontal.MulScalar(s).Add(c.vertical.MulScalar(t)).Sub(c.origin)).Sub(offset),
	}
}

func degreesToRadians(degrees math.Real) math.Real {
	return degrees * math.Pi / 180.0
}
