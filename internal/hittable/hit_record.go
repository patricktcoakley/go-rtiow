package hittable

import (
	"github.com/patricktcoakley/go-rtiow/internal/geometry"
	"github.com/patricktcoakley/go-rtiow/internal/math"
)

type HitRecord struct {
	Point     geometry.Vec3
	Normal    geometry.Vec3
	Material  Scatterable
	T         math.Real
	FrontFace bool
}

func (hr *HitRecord) SetFaceNormal(r geometry.Ray, outwardNormal geometry.Vec3) {
	hr.FrontFace = geometry.Dot(r.Direction, outwardNormal) < 0
	if hr.FrontFace {
		hr.Normal = outwardNormal
	} else {
		hr.Normal = outwardNormal.Neg()
	}
}
