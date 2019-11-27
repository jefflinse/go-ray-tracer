package rt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPattern_AtObject(t *testing.T) {
	// with object transformation
	s := NewSphere()
	s.Transform = NewScaling(2, 2, 2)
	p := NewPatternProps()
	c := p.atObject(s, NewPoint(2, 3, 4), func(patternPoint Tuple) Color {
		return NewColor(patternPoint.X(), patternPoint.Y(), patternPoint.Z())
	})
	assert.True(t, c.Equals(NewColor(1, 1.5, 2)))

	// with pattern transformation
	s = NewSphere()
	p = NewPatternProps()
	p.Transform = NewScaling(2, 2, 2)
	c = p.atObject(s, NewPoint(2, 3, 4), func(patternPoint Tuple) Color {
		return NewColor(patternPoint.X(), patternPoint.Y(), patternPoint.Z())
	})
	assert.Equal(t, NewColor(1, 1.5, 2), c)
	assert.True(t, c.Equals(NewColor(1, 1.5, 2)))

	// with both object and pattern transformations
	s = NewSphere()
	s.Transform = NewScaling(2, 2, 2)
	p = NewPatternProps()
	p.Transform = NewTranslation(.5, 1, 1.5)
	c = p.atObject(s, NewPoint(2.5, 3, 3.5), func(patternPoint Tuple) Color {
		return NewColor(patternPoint.X(), patternPoint.Y(), patternPoint.Z())
	})
	assert.True(t, c.Equals(NewColor(.75, .5, .25)))
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
