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
	atA(parentPatternPoint Tuple) Color
	atB(parentPatternPoint Tuple) Color
}

// PatternProps contains properties common to all patterns.
type PatternProps struct {
	A         Pattern
	B         Pattern
	Transform Transformation
	p         Pattern
}

// NewPatternProps creates a new PatternProps.
func NewPatternProps(a Pattern, b Pattern) PatternProps {
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

// Returns the color at subpattern A by first applying its transformation.
func (props *PatternProps) atA(parentPatternPoint Tuple) Color {
	return props.A.At(props.A.GetTransform().Inverse().ApplyTo(parentPatternPoint))
}

// Returns the color at subpattern B by first applying its transformation.
func (props *PatternProps) atB(parentPatternPoint Tuple) Color {
	return props.B.At(props.B.GetTransform().Inverse().ApplyTo(parentPatternPoint))
}

// A SolidPattern is just a single color.
type SolidPattern struct {
	PatternProps
	color Color
}

// NewSolidPattern creates a new SolidPattern.
func NewSolidPattern(color Color) *SolidPattern {
	pattern := &SolidPattern{NewPatternProps(nil, nil), color}
	pattern.p = pattern
	return pattern
}

// At returns the pattern color at the given point.
func (p *SolidPattern) At(point Tuple) Color {
	return p.color
}

// A BlendedPattern is blended combination of two other patterns.
type BlendedPattern struct {
	PatternProps
}

// NewBlendedPattern creates a new BlendedPatter.
func NewBlendedPattern(a Pattern, b Pattern) *BlendedPattern {
	pattern := &BlendedPattern{NewPatternProps(a, b)}
	pattern.p = pattern
	return pattern
}

// At returns the pattern color at the given point.
func (p *BlendedPattern) At(point Tuple) Color {
	colorA := p.atA(point)
	colorB := p.atB(point)
	return colorA.AverageBlend(colorB)
}

// A StripePattern is a pattern of colors alternates in the X axis.
type StripePattern struct {
	PatternProps
}

// NewStripePattern creates a new StripePattern.
func NewStripePattern(a Pattern, b Pattern) *StripePattern {
	pattern := &StripePattern{NewPatternProps(a, b)}
	pattern.p = pattern
	return pattern
}

// At returns the pattern color at the given point.
func (p *StripePattern) At(point Tuple) Color {
	if int(math.Floor(point.X()))%2 == 0 {
		return p.atA(point)
	}

	return p.atB(point)
}

// A GradientPattern is a pattern that fades from one color to another.
type GradientPattern struct {
	PatternProps
}

// NewGradientPattern creates a new GradientPattern.
func NewGradientPattern(a Pattern, b Pattern) *GradientPattern {
	pattern := &GradientPattern{NewPatternProps(a, b)}
	pattern.p = pattern
	return pattern
}

// At returns the pattern color at the given point.
func (p *GradientPattern) At(point Tuple) Color {
	atA := p.atA(point)
	atB := p.atB(point)
	distance := atB.Subtract(atA)
	fraction := point.X() - math.Floor(point.X())
	return atA.Add(distance.Multiply(fraction))
}

// A RingPattern is a pattern of alternating rings of color.
type RingPattern struct {
	PatternProps
}

// NewRingPattern creates a new RingPattern.
func NewRingPattern(a Pattern, b Pattern) *RingPattern {
	pattern := &RingPattern{NewPatternProps(a, b)}
	pattern.p = pattern
	return pattern
}

// At returns the pattern color at the given point.
func (p *RingPattern) At(point Tuple) Color {
	if int(math.Floor(math.Sqrt(math.Pow(point.X(), 2)+math.Pow(point.Z(), 2))))%2 == 0 {
		return p.atA(point)
	}

	return p.atB(point)
}

// A CheckerPattern is a pattern of alternating colors in all dimensions.
type CheckerPattern struct {
	PatternProps
}

// NewCheckerPattern creates a new RingPattern.
func NewCheckerPattern(a Pattern, b Pattern) *CheckerPattern {
	pattern := &CheckerPattern{NewPatternProps(a, b)}
	pattern.p = pattern
	return pattern
}

// At returns the pattern color at the given point.
func (p *CheckerPattern) At(point Tuple) Color {
	if (int(math.Floor(point.X()))+int(math.Floor(point.Y()))+int(math.Floor(point.Z())))%2 == 0 {
		subpatternPoint := p.A.GetTransform().Inverse().ApplyTo(point)
		return p.A.At(subpatternPoint)
	}

	subpatternPoint := p.B.GetTransform().Inverse().ApplyTo(point)
	return p.B.At(subpatternPoint)
}
