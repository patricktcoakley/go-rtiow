package hittable

import (
	"math"
	"math/rand"

	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

type Dielectric struct {
	Ir float64
}

func NewDielectric(ir float64) *Dielectric {
	return &Dielectric{ir}
}

func (d *Dielectric) Scatter(rIn ray.Ray, hr *HitRecord, attenuation *vec3.Vec3, scattered *ray.Ray) bool {
	*attenuation = vec3.Vec3{1, 1, 1}
	refractionRatio := d.Ir
	if hr.FrontFace {
		refractionRatio = 1 / d.Ir
	}

	unitDirection := rIn.Direction.ToUnit()
	cosTheta := math.Min(vec3.Dot(unitDirection.Neg(), hr.Normal), 1.)
	sinTheta := math.Sqrt(1 - cosTheta*cosTheta)
	cannotRefract := refractionRatio*sinTheta > 1.
	var direction vec3.Vec3
	if cannotRefract || reflectance(cosTheta, refractionRatio) > rand.Float64() {
		direction = vec3.Reflect(unitDirection, hr.Normal)
	} else {
		direction = vec3.Refract(unitDirection, hr.Normal, refractionRatio)
	}

	*scattered = ray.Ray{Origin: hr.Point, Direction: direction}
	return true
}

func reflectance(cos, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 *= r0
	return r0 + (1-r0)*math.Pow((1-cos), 5)
}
