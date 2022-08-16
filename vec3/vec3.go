package vec3

import "math"

type Vec3 [3]float64

func (v *Vec3) Neg() Vec3 {
	return Vec3{-v[0], -v[1], -v[2]}
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(Dot(v, v))
}

func (v *Vec3) ToUnit() Vec3 {
	return DivScalar(v, v.Length())
}

func Add(lhs, rhs *Vec3) Vec3 {
	return Vec3{lhs[0] + rhs[0], lhs[1] + rhs[1], lhs[2] + rhs[2]}
}

func Sub(lhs, rhs *Vec3) Vec3 {
	return Vec3{lhs[0] - rhs[0], lhs[1] - rhs[1], lhs[2] - rhs[2]}
}

func Mul(lhs, rhs *Vec3) Vec3 {
	return Vec3{lhs[0] * rhs[0], lhs[1] * rhs[1], lhs[2] * rhs[2]}
}

func MulScalar(lhs *Vec3, rhs float64) Vec3 {
	return Vec3{lhs[0] * rhs, lhs[1] * rhs, lhs[2] * rhs}
}

func Div(lhs, rhs *Vec3) Vec3 {
	return Vec3{lhs[0] / rhs[0], lhs[1] / rhs[1], lhs[2] / rhs[2]}
}

func DivScalar(lhs *Vec3, rhs float64) Vec3 {
	rhs = 1.0 / rhs
	return Vec3{lhs[0] * rhs, lhs[1] * rhs, lhs[2] * rhs}
}

func Dot(lhs, rhs *Vec3) float64 {
	return lhs[0]*rhs[0] + lhs[1]*rhs[1] + lhs[2]*rhs[2]
}

func Cross(lhs, rhs *Vec3) Vec3 {
	return Vec3{
		lhs[1]*rhs[2] - lhs[2]*rhs[1],
		lhs[2]*rhs[0] - lhs[0]*rhs[2],
		lhs[0]*rhs[1] - lhs[1]*rhs[0],
	}
}
