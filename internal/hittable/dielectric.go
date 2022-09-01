package hittable

import (
	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

type Dielectric struct {
	Ir float64
}

func NewDieletric(ir float64) *Dielectric {
	return &Dielectric{ir}
}

func (d *Dielectric) Scatter(rIn ray.Ray, hr HitRecord, attenuation *vec3.Vec3, scattered *ray.Ray) bool {
	*attenuation = vec3.Vec3{1, 1, 1}
	var refractionRatio float64
	if hr.FrontFace {
		refractionRatio = 1.0 / d.Ir
	} else {
		refractionRatio = d.Ir
	}
	
	unitDirection := rIn.Direction.ToUnit()
	refracted := vec3.Refract(unitDirection, hr.Normal, refractionRatio)
	*scattered = ray.Ray{Origin: hr.Point, Direction: refracted}
	return true
}
