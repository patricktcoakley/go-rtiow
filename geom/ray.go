package geom

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (r Ray) At(t float64) Vec3 {
	return Add(r.Origin, MulScalar(r.Direction, t))
}
