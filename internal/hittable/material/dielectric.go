package material

import (
	"github.com/patricktcoakley/go-rtiow/internal/geometry"
	"github.com/patricktcoakley/go-rtiow/internal/hittable"
	"github.com/patricktcoakley/go-rtiow/internal/math"
)

type Dielectric struct {
	Ir math.Real
}

func NewDielectric(ir math.Real) *Dielectric {
	return &Dielectric{ir}
}

func (d *Dielectric) Scatter(rIn geometry.Ray, hr *hittable.HitRecord, attenuation *geometry.Vec3, scattered *geometry.Ray) bool {
	*attenuation = geometry.Vec3{X: 1, Y: 1, Z: 1}
	refractionRatio := d.Ir
	if hr.FrontFace {
		refractionRatio = 1 / d.Ir
	}

	unitDirection := rIn.Direction.ToUnit()
	cosTheta := math.Min(geometry.Dot(unitDirection.Neg(), hr.Normal), 1.)
	sinTheta := math.Sqrt(1 - cosTheta*cosTheta)
	cannotRefract := refractionRatio*sinTheta > 1.
	var direction geometry.Vec3
	if cannotRefract || reflectance(cosTheta, refractionRatio) > math.Random() {
		direction = geometry.Reflect(unitDirection, hr.Normal)
	} else {
		direction = geometry.Refract(unitDirection, hr.Normal, refractionRatio)
	}

	*scattered = geometry.Ray{Origin: hr.Point, Direction: direction}
	return true
}

func reflectance(cos, refIdx math.Real) math.Real {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 *= r0
	return r0 + (1-r0)*math.Pow(1-cos, 5)
}
