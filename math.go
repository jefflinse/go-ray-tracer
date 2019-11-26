package rt

import (
	"math"
)

// EPSILON is the amount by which two floats must differ to be considered different.
const EPSILON = .00001

func clamp(value float64, min float64, max float64) float64 {
	return math.Min(math.Max(value, min), max)
}

func eq(a float64, b float64) bool {
	return math.Abs(a-b) < EPSILON
}
