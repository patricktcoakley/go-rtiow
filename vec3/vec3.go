package vec3

import "math"

type Vec3 [3]float64
type Color [3]uint8

func (v Vec3) Neg() Vec3 {
	return Vec3{-v[0], -v[1], -v[2]}
}

func (v Vec3) Length() float64 {
	return math.Sqrt(Dot(v, v))
}

func (v Vec3) LengthSquared() float64 {
	return v.Length() * v.Length()
}

func (v Vec3) ToUnit() Vec3 {
	return DivScalar(v, v.Length())
}

func (v Vec3) ToRGB() Color {
	return Color{
		uint8(v[0] * 255.99),
		uint8(v[1] * 255.99),
		uint8(v[2] * 255.99),
	}
}

func Add(lhs, rhs Vec3) Vec3 {
	return Vec3{lhs[0] + rhs[0], lhs[1] + rhs[1], lhs[2] + rhs[2]}
}

func (v Vec3) Add(rhs Vec3) Vec3 {
	return Add(v, rhs)
}

func Sub(lhs, rhs Vec3) Vec3 {
	return Vec3{lhs[0] - rhs[0], lhs[1] - rhs[1], lhs[2] - rhs[2]}
}

func (v Vec3) Sub(rhs Vec3) Vec3 {
	return Sub(v, rhs)
}

func Mul(lhs, rhs Vec3) Vec3 {
	return Vec3{lhs[0] * rhs[0], lhs[1] * rhs[1], lhs[2] * rhs[2]}
}

func (v Vec3) Mul(rhs Vec3) Vec3 {
	return Mul(v, rhs)
}

func MulScalar(lhs Vec3, rhs float64) Vec3 {
	return Vec3{lhs[0] * rhs, lhs[1] * rhs, lhs[2] * rhs}
}

func (v Vec3) MulScalar(rhs float64) Vec3 {
	return MulScalar(v, rhs)
}

func Div(lhs, rhs Vec3) Vec3 {
	return Vec3{lhs[0] / rhs[0], lhs[1] / rhs[1], lhs[2] / rhs[2]}
}

func (v Vec3) Div(rhs Vec3) Vec3 {
	return Div(v, rhs)
}

func DivScalar(lhs Vec3, rhs float64) Vec3 {
	rhs = 1.0 / rhs
	return Vec3{lhs[0] * rhs, lhs[1] * rhs, lhs[2] * rhs}
}

func (v Vec3) DivScalar(rhs float64) Vec3 {
	return DivScalar(v, rhs)
}

func Dot(lhs, rhs Vec3) float64 {
	return lhs[0]*rhs[0] + lhs[1]*rhs[1] + lhs[2]*rhs[2]
}

func (v Vec3) Dot(rhs Vec3) float64 {
	return Dot(v, rhs)
}

func Cross(lhs, rhs Vec3) Vec3 {
	return Vec3{
		lhs[1]*rhs[2] - lhs[2]*rhs[1],
		lhs[2]*rhs[0] - lhs[0]*rhs[2],
		lhs[0]*rhs[1] - lhs[1]*rhs[0],
	}
}

func (v Vec3) Cross(rhs Vec3) Vec3 {
	return Cross(v, rhs)
}
