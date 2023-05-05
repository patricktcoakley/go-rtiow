package hittable

import (
	"github.com/patricktcoakley/go-rtiow/internal/math"
	"github.com/patricktcoakley/go-rtiow/internal/ray"
	"github.com/patricktcoakley/go-rtiow/internal/vec3"
)

type HitRecord struct {
	Point     vec3.Vec3
	Normal    vec3.Vec3
	Material  Material
	T         math.Real
	FrontFace bool
}

func (hr *HitRecord) SetFaceNormal(r ray.Ray, outwardNormal vec3.Vec3) {
	hr.FrontFace = vec3.Dot(r.Direction, outwardNormal) < 0
	if hr.FrontFace {
		hr.Normal = outwardNormal
	} else {
		hr.Normal = outwardNormal.Neg()
	}
}
