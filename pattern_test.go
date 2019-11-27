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

func TestStripePattern_At(t *testing.T) {
	// pattern is constant in Y
	p := NewStripePattern(white, black)
	assert.Equal(t, white, p.ColorAt(NewPoint(0, 0, 0)))
	assert.Equal(t, white, p.ColorAt(NewPoint(0, 1, 0)))
	assert.Equal(t, white, p.ColorAt(NewPoint(0, 2, 0)))

	// pattern is constant in Z
	p = NewStripePattern(white, black)
	assert.Equal(t, white, p.ColorAt(NewPoint(0, 0, 0)))
	assert.Equal(t, white, p.ColorAt(NewPoint(0, 0, 1)))
	assert.Equal(t, white, p.ColorAt(NewPoint(0, 0, 2)))

	// alternates in X
	p = NewStripePattern(white, black)
	assert.Equal(t, white, p.ColorAt(NewPoint(0, 0, 0)))
	assert.Equal(t, white, p.ColorAt(NewPoint(.9, 0, 0)))
	assert.Equal(t, black, p.ColorAt(NewPoint(1, 0, 0)))
	assert.Equal(t, black, p.ColorAt(NewPoint(-.1, 0, 0)))
	assert.Equal(t, black, p.ColorAt(NewPoint(-1, 0, 0)))
	assert.Equal(t, white, p.ColorAt(NewPoint(-1.1, 0, 0)))
}
