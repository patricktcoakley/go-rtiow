package vec3

import (
	"math"
	"math/rand"
)

type Vec3 [3]float64

func (v Vec3) Neg() Vec3 {
	return Vec3{-v[0], -v[1], -v[2]}
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) LengthSquared() float64 {
	return Dot(v, v)
}

func (v Vec3) ToUnit() Vec3 {
	return DivScalar(v, v.Length())
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

func (v Vec3) NearZero() bool {
	s := 1e-8
	return math.Abs(v[0]) < s && math.Abs(v[1]) < s && math.Abs(v[2]) < s
}

func Reflect(lhs, rhs Vec3) Vec3 {
	return Sub(
		lhs,
		MulScalar(MulScalar(rhs, Dot(lhs, rhs)), 2.),
	)
}

func Refract(uv, n Vec3, etai_over_etat float64) Vec3 {
	cosTheta := math.Min(Dot(uv.Neg(), n), 1.0)
	rOutPerp := MulScalar(
		Add(uv, MulScalar(n, cosTheta)),
		etai_over_etat,
	)

	rOutParallel := MulScalar(
		n,
		-math.Sqrt(math.Abs(1.0-rOutPerp.LengthSquared())),
	)
	return Add(rOutPerp, rOutParallel)
}

func NewRandomVec3() Vec3 {
	return Vec3{rand.Float64(), rand.Float64(), rand.Float64()}
}

func NewRandomRangeVec3(min, max float64) Vec3 {
	x := (min) + (max-min)*rand.Float64()
	y := (min) + (max-min)*rand.Float64()
	z := (min) + (max-min)*rand.Float64()

	return Vec3{x, y, z}
}

func RandomInUnitSphere() Vec3 {
	p := NewRandomRangeVec3(-1, 1)
	for p.LengthSquared() >= 1 {
		p = NewRandomRangeVec3(-1, 1)
	}

	return p
}

func NewRandomUnitVector() Vec3 {
	return RandomInUnitSphere().ToUnit()
}

func RandomInHempishphere(normal Vec3) Vec3 {
	inUnitSphere := RandomInUnitSphere()
	if Dot(inUnitSphere, normal) > 0. {
		return inUnitSphere
	}

	return inUnitSphere.Neg()
}
