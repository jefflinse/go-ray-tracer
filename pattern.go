package rt

import (
	"math"
)

// Color constants
var black = NewColor(0, 0, 0)
var white = NewColor(1, 1, 1)

// A Pattern is a pattern that can be applied to a surface.
type Pattern interface {
	At(point Tuple) Color
	AtObject(object Shape, point Tuple) Color
}

// A StripePattern is a pattern of colors alternates in the X axis.
type StripePattern struct {
	A         Color
	B         Color
	Transform Transformation
}

// NewStripePattern creates a new StripePattern.
func NewStripePattern(a Color, b Color) *StripePattern {
	return &StripePattern{a, b, NewTransform()}
}

// At returns the pattern color at the given point.
func (p *StripePattern) At(point Tuple) Color {
	if int(math.Floor(point.X()))%2 == 0 {
		return p.A
	}

	return p.B
}

// AtObject returns the pattern color on the specified object at the specified point.
func (p *StripePattern) AtObject(object Shape, point Tuple) Color {
	localPoint := object.GetTransform().Inverse().ApplyTo(point)
	patternPoint := p.Transform.Inverse().ApplyTo(localPoint)
	return p.At(patternPoint)
}
