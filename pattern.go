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

// PatternProps contains properties common to all patterns.
type PatternProps struct {
	Transform Transformation
}

// NewPatternProps creates a new PatternProps.
func NewPatternProps() PatternProps {
	return PatternProps{
		Transform: NewTransform(),
	}
}

// GetTransform returns the pattern's transformation.
func (p *PatternProps) GetTransform() Transformation {
	return p.Transform
}

func (p *PatternProps) atObject(object Shape, worldPoint Tuple, patternAtObjectFn func(patternPoint Tuple) Color) Color {
	localPoint := object.GetTransform().Inverse().ApplyTo(worldPoint)
	patternPoint := p.GetTransform().Inverse().ApplyTo(localPoint)
	return patternAtObjectFn(patternPoint)
}

// A StripePattern is a pattern of colors alternates in the X axis.
type StripePattern struct {
	PatternProps
	A Color
	B Color
}

// NewStripePattern creates a new StripePattern.
func NewStripePattern(a Color, b Color) *StripePattern {
	return &StripePattern{
		PatternProps: NewPatternProps(),
		A:            a,
		B:            b,
	}
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
	return p.atObject(object, point, func(patternPoint Tuple) Color {
		return p.At(patternPoint)
	})
}

// A GradientPattern is a pattern that fades from one color to another.
type GradientPattern struct {
	PatternProps
	A Color
	B Color
}

// NewGradientPattern creates a new GradientPattern.
func NewGradientPattern(a Color, b Color) *GradientPattern {
	return &GradientPattern{
		PatternProps: NewPatternProps(),
		A:            a,
		B:            b,
	}
}

// At returns the pattern color at the given point.
func (p *GradientPattern) At(point Tuple) Color {
	distance := p.B.Subtract(p.A)
	fraction := point.X() - math.Floor(point.X())
	return p.A.Add(distance.Multiply(fraction))
}

// AtObject returns the pattern color on the specified object at the specified point.
func (p *GradientPattern) AtObject(object Shape, point Tuple) Color {
	return p.atObject(object, point, func(patternPoint Tuple) Color {
		return p.At(patternPoint)
	})
}
