package hittable

import (
	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

type Metal struct {
	Albedo    vec3.Vec3
	fuzziness float64
}

func NewMetal(r, g, b, fuzziness float64) *Metal {
	if fuzziness >= 1 {
		fuzziness = 1
	}
	return &Metal{vec3.Vec3{r, g, b}, fuzziness}
}

func (m *Metal) Scatter(rIn ray.Ray, hr HitRecord, attenuation *vec3.Vec3, scattered *ray.Ray) bool {
	reflected := vec3.Reflect(rIn.Direction.ToUnit(), hr.Normal)
	*scattered = ray.Ray{Origin: hr.Point, Direction: reflected.Add(vec3.MulScalar(vec3.RandomInUnitSphere(), m.fuzziness))}
	*attenuation = m.Albedo
	return vec3.Dot(scattered.Direction, hr.Normal) > 0
}
