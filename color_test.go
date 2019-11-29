package rt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewColor(t *testing.T) {
	c1 := NewColor(-.5, .4, 1.7)
	assert.Equal(t, -.5, c1.Red())
	assert.Equal(t, .4, c1.Green())
	assert.Equal(t, 1.7, c1.Blue())
}

func TestColor_Equals(t *testing.T) {
	c1 := NewColor(.9, .6, .75)
	c2 := NewColor(.9, .6, .75)
	assert.True(t, c1.Equals(c2))
	c2 = NewColor(.7, .1, .25)
	assert.False(t, c1.Equals(c2))
}

func TestColor_Add(t *testing.T) {
	c1 := NewColor(.9, .6, .75)
	c2 := NewColor(.7, .1, .25)
	assert.Equal(t, NewColor(1.6, .7, 1), c1.Add(c2))
}

func TestColor_Subtract(t *testing.T) {
	c1 := NewColor(.9, .6, .75)
	c2 := NewColor(.7, .1, .25)
	assert.True(t, c1.Subtract(c2).Equals(NewColor(.2, .5, .5)))
}

func TestColor_Multiple(t *testing.T) {
	c1 := NewColor(.2, .3, .4)
	assert.True(t, c1.Multiply(2).Equals(NewColor(.4, .6, .8)))
}

func TestColor_AverageBlend(t *testing.T) {
	c1 := NewColor(1, .2, .4)
	c2 := NewColor(1, .4, .1)
	assert.True(t, c1.AverageBlend(c2).Equals(NewColor(1, .3, .25)))
}

func TestColor_HadamardBlend(t *testing.T) {
	c1 := NewColor(1, .2, .4)
	c2 := NewColor(.9, 1, .1)
	assert.True(t, c1.HadamardBlend(c2).Equals(NewColor(.9, .2, .04)))
}

func TestColor_ToPPM(t *testing.T) {
	var c Color
	assert.Equal(t, "0 0 0", c.ToPPM())
	c = NewColor(0, 0, 0)
	assert.Equal(t, "0 0 0", c.ToPPM())
	c = NewColor(1, 1, 1)
	assert.Equal(t, "255 255 255", c.ToPPM())
	c = NewColor(.5, .3, .1)
	assert.Equal(t, "128 77 26", c.ToPPM())
}
