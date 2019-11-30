package rt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var solidBlack = NewSolidPattern(black)
var solidWhite = NewSolidPattern(white)

type testPattern struct {
	PatternProps
}

func newTestPattern(a Pattern, b Pattern) *testPattern {
	pattern := &testPattern{NewPatternProps(a, b)}
	pattern.p = pattern
	return pattern
}

func (tp *testPattern) At(point Tuple) Color {
	return NewColor(point.X(), point.Y(), point.Z())
}

func TestNewPatternProps(t *testing.T) {
	p := NewPatternProps(solidWhite, solidBlack)
	assert.Equal(t, solidWhite, p.A)
	assert.Equal(t, solidBlack, p.B)
	assert.Equal(t, NewTransform(), p.Transform)
}

func TestPatternProps_AtObject(t *testing.T) {
	// without any transformations
	o := NewSphere()
	p := newTestPattern(solidWhite, solidBlack)
	c := p.AtObject(o, NewPoint(2, 3, 4))
	assert.Equal(t, NewColor(2, 3, 4), c)

	// with object transformation
	o = NewSphere()
	o.Transform = NewScaling(2, 2, 2)
	p = newTestPattern(solidWhite, solidBlack)
	c = p.AtObject(o, NewPoint(2, 3, 4))
	assert.Equal(t, NewColor(1, 1.5, 2), c)

	// with pattern transformation
	o = NewSphere()
	p = newTestPattern(solidWhite, solidBlack)
	p.SetTransform(NewScaling(2, 2, 2))
	c = p.AtObject(o, NewPoint(2, 3, 4))
	assert.Equal(t, NewColor(1, 1.5, 2), c)

	// with both object and pattern transformations
	o = NewSphere()
	o.Transform = NewScaling(2, 2, 2)
	p = newTestPattern(solidWhite, solidBlack)
	p.SetTransform(NewTranslation(.5, 1, 1.5))
	c = p.AtObject(o, NewPoint(2.5, 3, 3.5))
	assert.Equal(t, NewColor(.75, .5, .25), c)
}

func TestBlendedPattern_AtObject(t *testing.T) {
	// without any transformations
	o := NewSphere()
	p := NewBlendedPattern(newTestPattern(solidWhite, solidBlack), newTestPattern(solidWhite, solidBlack))
	c := p.AtObject(o, NewPoint(2, 3, 4))
	assert.Equal(t, NewColor(2, 3, 4), c)

	// with object transformation
	o = NewSphere()
	o.Transform = NewScaling(2, 2, 2)
	p = NewBlendedPattern(newTestPattern(solidWhite, solidBlack), newTestPattern(solidWhite, solidBlack))
	c = p.AtObject(o, NewPoint(2, 3, 4))
	assert.Equal(t, NewColor(1, 1.5, 2), c)

	// with pattern transformation
	o = NewSphere()
	p = NewBlendedPattern(newTestPattern(solidWhite, solidBlack), newTestPattern(solidWhite, solidBlack))
	p.SetTransform(NewScaling(2, 2, 2))
	c = p.AtObject(o, NewPoint(2, 3, 4))
	assert.Equal(t, NewColor(2, 3, 4), c)

	// with both object and pattern transformations
	o = NewSphere()
	o.Transform = NewScaling(2, 2, 2)
	p = NewBlendedPattern(newTestPattern(solidWhite, solidBlack), newTestPattern(solidWhite, solidBlack))
	p.SetTransform(NewTranslation(.5, 1, 1.5))
	c = p.AtObject(o, NewPoint(2.5, 3, 3.5))
	assert.Equal(t, NewColor(1.25, 1.5, 1.75), c)
}

func TestBlendedPattern_At(t *testing.T) {
	p := NewBlendedPattern(newTestPattern(solidWhite, solidBlack), newTestPattern(solidWhite, solidBlack))
	assert.Nil(t, p.At(NewPoint(0, 0, 0)))
}

func TestStripePattern_At(t *testing.T) {
	// pattern is constant in Y
	p := NewStripePattern(solidWhite, solidBlack)
	assert.Equal(t, white, p.At(NewPoint(0, 0, 0)))
	assert.Equal(t, white, p.At(NewPoint(0, 1, 0)))
	assert.Equal(t, white, p.At(NewPoint(0, 2, 0)))

	// pattern is constant in Z
	p = NewStripePattern(solidWhite, solidBlack)
	assert.Equal(t, white, p.At(NewPoint(0, 0, 0)))
	assert.Equal(t, white, p.At(NewPoint(0, 0, 1)))
	assert.Equal(t, white, p.At(NewPoint(0, 0, 2)))

	// alternates in X
	p = NewStripePattern(solidWhite, solidBlack)
	assert.Equal(t, white, p.At(NewPoint(0, 0, 0)))
	assert.Equal(t, white, p.At(NewPoint(.9, 0, 0)))
	assert.Equal(t, black, p.At(NewPoint(1, 0, 0)))
	assert.Equal(t, black, p.At(NewPoint(-.1, 0, 0)))
	assert.Equal(t, black, p.At(NewPoint(-1, 0, 0)))
	assert.Equal(t, white, p.At(NewPoint(-1.1, 0, 0)))
}

func TestGradientPattern_At(t *testing.T) {
	p := NewGradientPattern(solidWhite, solidBlack)
	assert.True(t, p.At(NewPoint(0, 0, 0)).Equals(white))
	assert.True(t, p.At(NewPoint(.25, 0, 0)).Equals(NewColor(.75, .75, .75)))
	assert.True(t, p.At(NewPoint(.5, 0, 0)).Equals(NewColor(.5, .5, .5)))
	assert.True(t, p.At(NewPoint(.75, 0, 0)).Equals(NewColor(.25, .25, .25)))
}

func TestRingPattern_At(t *testing.T) {
	p := NewRingPattern(solidWhite, solidBlack)
	assert.True(t, p.At(NewPoint(0, 0, 0)).Equals(white))
	assert.True(t, p.At(NewPoint(1, 0, 0)).Equals(black))
	assert.True(t, p.At(NewPoint(0, 0, 1)).Equals(black))
	assert.True(t, p.At(NewPoint(.708, 0, .708)).Equals(black))
}

func TestCheckerPattern_At(t *testing.T) {
	p := NewCheckerPattern(solidWhite, solidBlack)
	assert.True(t, p.At(NewPoint(0, 0, 0)).Equals(white))
	assert.True(t, p.At(NewPoint(.99, 0, 0)).Equals(white))
	assert.True(t, p.At(NewPoint(1.01, 0, 0)).Equals(black))
	assert.True(t, p.At(NewPoint(0, 0, 0)).Equals(white))
	assert.True(t, p.At(NewPoint(0, .99, 0)).Equals(white))
	assert.True(t, p.At(NewPoint(0, 1.01, 0)).Equals(black))
	assert.True(t, p.At(NewPoint(0, 0, 0)).Equals(white))
	assert.True(t, p.At(NewPoint(0, 0, .99)).Equals(white))
	assert.True(t, p.At(NewPoint(0, 0, 1.01)).Equals(black))
}
