package rt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStripePattern(t *testing.T) {
	p := NewStripePattern(white, black)
	assert.Equal(t, white, p.A)
	assert.Equal(t, black, p.B)
}

func TestStripePattern_ColorAt(t *testing.T) {
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

func TestStripePattern_ColorAtObject(t *testing.T) {
	// with object transformation
	s := NewSphere()
	s.Transform = NewScaling(2, 2, 2)
	p := NewStripePattern(white, black)
	c := p.AtObject(s, NewPoint(1.5, 0, 0))
	assert.Equal(t, white, c)

	// with pattern transformation
	s = NewSphere()
	p = NewStripePattern(white, black)
	p.Transform = NewScaling(2, 2, 2)
	c = p.AtObject(s, NewPoint(1.5, 0, 0))
	assert.Equal(t, white, c)

	// with both object and pattern transformations
	s = NewSphere()
	s.Transform = NewScaling(2, 2, 2)
	p = NewStripePattern(white, black)
	p.Transform = NewTranslation(.5, 0, 0)
	c = p.AtObject(s, NewPoint(2.5, 0, 0))
	assert.Equal(t, white, c)
}
