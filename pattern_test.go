package rt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testPattern struct {
	PatternProps
}

func newTestPattern(a Color, b Color) *testPattern {
	pattern := &testPattern{NewPatternProps(a, b)}
	pattern.p = pattern
	return pattern
}

func (tp *testPattern) At(point Tuple) Color {
	return NewColor(point.X(), point.Y(), point.Z())
}

func TestPatternProps_AtObject(t *testing.T) {
	// without any transformations
	o := NewSphere()
	p := newTestPattern(white, black)
	c := p.AtObject(o, NewPoint(2, 3, 4))
	assert.Equal(t, NewColor(2, 3, 4), c)

	// with object transformation
	o = NewSphere()
	o.Transform = NewScaling(2, 2, 2)
	p = newTestPattern(white, black)
	c = p.AtObject(o, NewPoint(2, 3, 4))
	assert.Equal(t, NewColor(1, 1.5, 2), c)

	// with pattern transformation
	o = NewSphere()
	p = newTestPattern(white, black)
	p.SetTransform(NewScaling(2, 2, 2))
	c = p.AtObject(o, NewPoint(2, 3, 4))
	assert.Equal(t, NewColor(1, 1.5, 2), c)

	// with both object and pattern transformations
	o = NewSphere()
	o.Transform = NewScaling(2, 2, 2)
	p = newTestPattern(white, black)
	p.SetTransform(NewTranslation(.5, 1, 1.5))
	c = p.AtObject(o, NewPoint(2.5, 3, 3.5))
	assert.Equal(t, NewColor(.75, .5, .25), c)
}

func TestNewStripePattern(t *testing.T) {
	p := NewStripePattern(white, black)
	assert.Equal(t, white, p.A)
	assert.Equal(t, black, p.B)
}

func TestStripePattern_At(t *testing.T) {
	// pattern is constant in Y
	p := NewStripePattern(white, black)
	assert.Equal(t, white, p.At(NewPoint(0, 0, 0)))
	assert.Equal(t, white, p.At(NewPoint(0, 1, 0)))
	assert.Equal(t, white, p.At(NewPoint(0, 2, 0)))

	// pattern is constant in Z
	p = NewStripePattern(white, black)
	assert.Equal(t, white, p.At(NewPoint(0, 0, 0)))
	assert.Equal(t, white, p.At(NewPoint(0, 0, 1)))
	assert.Equal(t, white, p.At(NewPoint(0, 0, 2)))

	// alternates in X
	p = NewStripePattern(white, black)
	assert.Equal(t, white, p.At(NewPoint(0, 0, 0)))
	assert.Equal(t, white, p.At(NewPoint(.9, 0, 0)))
	assert.Equal(t, black, p.At(NewPoint(1, 0, 0)))
	assert.Equal(t, black, p.At(NewPoint(-.1, 0, 0)))
	assert.Equal(t, black, p.At(NewPoint(-1, 0, 0)))
	assert.Equal(t, white, p.At(NewPoint(-1.1, 0, 0)))
}

func TestGradientPattern_At(t *testing.T) {
	p := NewGradientPattern(white, black)
	assert.True(t, p.At(NewPoint(0, 0, 0)).Equals(white))
	assert.True(t, p.At(NewPoint(.25, 0, 0)).Equals(NewColor(.75, .75, .75)))
	assert.True(t, p.At(NewPoint(.5, 0, 0)).Equals(NewColor(.5, .5, .5)))
	assert.True(t, p.At(NewPoint(.75, 0, 0)).Equals(NewColor(.25, .25, .25)))
}

func TestRingPattern_At(t *testing.T) {
	p := NewRingPattern(white, black)
	assert.True(t, p.At(NewPoint(0, 0, 0)).Equals(white))
	assert.True(t, p.At(NewPoint(1, 0, 0)).Equals(black))
	assert.True(t, p.At(NewPoint(0, 0, 1)).Equals(black))
	assert.True(t, p.At(NewPoint(.708, 0, .708)).Equals(black))
}
