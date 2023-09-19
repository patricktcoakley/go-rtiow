package geometry

import (
	"github.com/patricktcoakley/go-rtiow/internal/math"
)

type Vec3 struct {
	X, Y, Z math.Real
}

func (v Vec3) Neg() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

func (v Vec3) Length() math.Real {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) LengthSquared() math.Real {
	return Dot(v, v)
}

func (v Vec3) ToUnit() Vec3 {
	return DivScalar(v, v.Length())
}

func Add(lhs, rhs Vec3) Vec3 {
	return Vec3{lhs.X + rhs.X, lhs.Y + rhs.Y, lhs.Z + rhs.Z}
}

func (v Vec3) Add(rhs Vec3) Vec3 {
	return Add(v, rhs)
}

func Sub(lhs, rhs Vec3) Vec3 {
	return Vec3{lhs.X - rhs.X, lhs.Y - rhs.Y, lhs.Z - rhs.Z}
}

func (v Vec3) Sub(rhs Vec3) Vec3 {
	return Sub(v, rhs)
}

func Mul(lhs, rhs Vec3) Vec3 {
	return Vec3{lhs.X * rhs.X, lhs.Y * rhs.Y, lhs.Z * rhs.Z}
}

func (v Vec3) Mul(rhs Vec3) Vec3 {
	return Mul(v, rhs)
}

func MulScalar(lhs Vec3, rhs math.Real) Vec3 {
	return Vec3{lhs.X * rhs, lhs.Y * rhs, lhs.Z * rhs}
}

func (v Vec3) MulScalar(rhs math.Real) Vec3 {
	return MulScalar(v, rhs)
}

func Div(lhs, rhs Vec3) Vec3 {
	return Vec3{lhs.X / rhs.X, lhs.Y / rhs.Y, lhs.Z / rhs.Z}
}

func (v Vec3) Div(rhs Vec3) Vec3 {
	return Div(v, rhs)
}

func DivScalar(lhs Vec3, rhs math.Real) Vec3 {
	rhs = 1.0 / rhs
	return Vec3{lhs.X * rhs, lhs.Y * rhs, lhs.Z * rhs}
}

func (v Vec3) DivScalar(rhs math.Real) Vec3 {
	return DivScalar(v, rhs)
}

func Dot(lhs, rhs Vec3) math.Real {
	return lhs.X*rhs.X + lhs.Y*rhs.Y + lhs.Z*rhs.Z
}

func (v Vec3) Dot(rhs Vec3) math.Real {
	return Dot(v, rhs)
}

func Cross(lhs, rhs Vec3) Vec3 {
	return Vec3{
		lhs.Y*rhs.Z - lhs.Z*rhs.Y,
		lhs.Z*rhs.X - lhs.X*rhs.Z,
		lhs.X*rhs.Y - lhs.Y*rhs.X,
	}
}

func (v Vec3) Cross(rhs Vec3) Vec3 {
	return Cross(v, rhs)
}

func (v Vec3) NearZero() bool {
	s := math.Real(1e-8)
	return math.Abs(v.X) < s && math.Abs(v.Y) < s && math.Abs(v.Z) < s
}

func Reflect(lhs, rhs Vec3) Vec3 {
	return Sub(
		lhs,
		MulScalar(MulScalar(rhs, Dot(lhs, rhs)), 2.),
	)
}

func Refract(uv, n Vec3, etaiOverEtat math.Real) Vec3 {
	cosTheta := math.Min(Dot(uv.Neg(), n), 1.0)
	rOutPerp := MulScalar(
		Add(uv, MulScalar(n, cosTheta)),
		etaiOverEtat,
	)

	rOutParallel := MulScalar(
		n,
		-math.Sqrt(math.Abs(1.0-rOutPerp.LengthSquared())),
	)
	return Add(rOutPerp, rOutParallel)
}

func NewRandomVec3() Vec3 {
	return Vec3{math.Random(), math.Random(), math.Random()}
}

func NewRandomRangeVec3(min, max math.Real) Vec3 {
	x := (min) + (max-min)*math.Random()
	y := (min) + (max-min)*math.Random()
	z := (min) + (max-min)*math.Random()

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

func RandomInHemisphere(normal Vec3) Vec3 {
	inUnitSphere := RandomInUnitSphere()
	if Dot(inUnitSphere, normal) > 0. {
		return inUnitSphere
	}

	return inUnitSphere.Neg()
}

func RandomInUnitDisk() Vec3 {
	for {
		p := Vec3{math.Random(), math.Random(), 0}
		if p.LengthSquared() < 1 {
			return p
		}
	}
}
