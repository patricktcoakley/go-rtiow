package math

import (
	"math"
	"math/rand/v2"
)

type Real float32

const (
	MinReal Real = math.SmallestNonzeroFloat32
	MaxReal Real = math.MaxFloat32
	Pi      Real = math.Pi
)

func Random() Real {
	return Real(rand.Float32())
}

func Sqrt(t Real) Real {
	return Real(math.Sqrt(float64(t)))
}

func Abs(t Real) Real {
	return Real(math.Abs(float64(t)))
}

func Pow(x, y Real) Real {
	return Real(math.Pow(float64(x), float64(y)))
}

func Tan(x Real) Real {
	return Real(math.Tan(float64(x)))
}

func Clamp(min, max, x uint8) uint8 {
	if x < min {
		return min
	}

	if x > max {
		return max
	}

	return x
}
