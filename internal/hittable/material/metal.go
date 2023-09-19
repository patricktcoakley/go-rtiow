package material

import (
	"github.com/patricktcoakley/go-rtiow/internal/geometry"
	"github.com/patricktcoakley/go-rtiow/internal/hittable"
	"github.com/patricktcoakley/go-rtiow/internal/math"
)

type Metal struct {
	Albedo    geometry.Vec3
	fuzziness math.Real
}

func NewMetal(r, g, b, fuzziness math.Real) *Metal {
	if fuzziness >= 1 {
		fuzziness = 1
	}
	return &Metal{geometry.Vec3{X: r, Y: g, Z: b}, fuzziness}
}

func (m *Metal) Scatter(rIn geometry.Ray, hr *hittable.HitRecord, attenuation *geometry.Vec3, scattered *geometry.Ray) bool {
	reflected := geometry.Reflect(rIn.Direction.ToUnit(), hr.Normal)
	*scattered = geometry.Ray{Origin: hr.Point, Direction: reflected.Add(geometry.MulScalar(geometry.RandomInUnitSphere(), m.fuzziness))}
	*attenuation = m.Albedo
	return geometry.Dot(scattered.Direction, hr.Normal) > 0
}
