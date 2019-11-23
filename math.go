package raytracer

import (
	"math"
)

func eq(a float64, b float64) bool {
	return math.Abs(a-b) < .00001
}
