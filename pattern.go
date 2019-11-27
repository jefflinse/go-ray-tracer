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
	GetTransform() Transformation
	SetTransform(transform Transformation)
}

// PatternProps contains properties common to all patterns.
type PatternProps struct {
	A         Color
	B         Color
	Transform Transformation
	p         Pattern
}

// NewPatternProps creates a new PatternProps.
func NewPatternProps(a Color, b Color) PatternProps {
	return PatternProps{
		A:         a,
		B:         b,
		Transform: NewTransform(),
	}
}

// GetTransform returns the pattern's transformation.
func (props *PatternProps) GetTransform() Transformation {
	return props.Transform
}

// SetTransform sets the pattern's transformation.
func (props *PatternProps) SetTransform(transform Transformation) {
	props.Transform = transform
}

// AtObject returns the pattern color on the specified object at the specified point.
func (props *PatternProps) AtObject(object Shape, worldPoint Tuple) Color {
	localPoint := object.GetTransform().Inverse().ApplyTo(worldPoint)
	patternPoint := props.p.GetTransform().Inverse().ApplyTo(localPoint)
	return props.p.At(patternPoint)
}

// A StripePattern is a pattern of colors alternates in the X axis.
type StripePattern struct {
	PatternProps
}

// NewStripePattern creates a new StripePattern.
func NewStripePattern(a Color, b Color) *StripePattern {
	pattern := &StripePattern{NewPatternProps(a, b)}
	pattern.p = pattern
	return pattern
}

// At returns the pattern color at the given point.
func (p *StripePattern) At(point Tuple) Color {
	if int(math.Floor(point.X()))%2 == 0 {
		return p.A
	}

	return p.B
}

// A GradientPattern is a pattern that fades from one color to another.
type GradientPattern struct {
	PatternProps
}

// NewGradientPattern creates a new GradientPattern.
func NewGradientPattern(a Color, b Color) *GradientPattern {
	pattern := &GradientPattern{NewPatternProps(a, b)}
	pattern.p = pattern
	return pattern
}

// At returns the pattern color at the given point.
func (p *GradientPattern) At(point Tuple) Color {
	distance := p.B.Subtract(p.A)
	fraction := point.X() - math.Floor(point.X())
	return p.A.Add(distance.Multiply(fraction))
}

// A RingPattern is a pattern of alternating rings of color.
type RingPattern struct {
	PatternProps
}

// NewRingPattern creates a new RingPattern.
func NewRingPattern(a Color, b Color) *RingPattern {
	pattern := &RingPattern{NewPatternProps(a, b)}
	pattern.p = pattern
	return pattern
}

// At returns the pattern color at the given point.
func (p *RingPattern) At(point Tuple) Color {
	if int(math.Floor(math.Sqrt(math.Pow(point.X(), 2)+math.Pow(point.Z(), 2))))%2 == 0 {
		return p.A
	}

	return p.B
}

// A CheckerPattern is a pattern of alternating colors in all dimensions.
type CheckerPattern struct {
	PatternProps
}

// NewCheckerPattern creates a new RingPattern.
func NewCheckerPattern(a Color, b Color) *CheckerPattern {
	pattern := &CheckerPattern{NewPatternProps(a, b)}
	pattern.p = pattern
	return pattern
}

// At returns the pattern color at the given point.
func (p *CheckerPattern) At(point Tuple) Color {
	if (int(math.Floor(point.X()))+int(math.Floor(point.Y()))+int(math.Floor(point.Z())))%2 == 0 {
		return p.A
	}

	return p.B
}