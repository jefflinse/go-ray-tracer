package rt

import (
	"math"
)

// Color constants
var black = NewColor(0, 0, 0)
var white = NewColor(1, 1, 1)

// A Pattern is a pattern that can be applied to a surface.
type Pattern interface {
	ColorAt(point Tuple) Color
}

// A StripePattern is a pattern of colors alternates in the X axis.
type StripePattern struct {
	A Color
	B Color
}

// NewStripePattern creates a new StripePattern.
func NewStripePattern(a Color, b Color) *StripePattern {
	return &StripePattern{a, b}
}

// ColorAt returns the pattern color at the given point.
func (p *StripePattern) ColorAt(point Tuple) Color {
	if int(math.Floor(point.X()))%2 == 0 {
		return p.A
	}

	return p.B
}
