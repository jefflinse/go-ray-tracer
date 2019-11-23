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

func TestColor_Add(t *testing.T) {
	c1 := NewColor(.9, .6, .75)
	c2 := NewColor(.7, .1, .25)
	assert.True(t, c1.Add(c2).Equals(NewColor(1.6, .7, 1)))
}

func TestColor_Subtract(t *testing.T) {
	c1 := NewColor(.9, .6, .75)
	c2 := NewColor(.7, .1, .25)
	assert.True(t, c1.Subtract(c2).Equals(NewColor(.2, .5, .5)))
}

func TestColor_Multiple(t *testing.T) {
	c1 := NewColor(.2, .3, .4)
	assert.Equal(t, NewColor(.4, .6, .8), c1.Multiply(2))
}

func TestColor_Blend(t *testing.T) {
	c1 := NewColor(1, .2, .4)
	c2 := NewColor(.9, 1, .1)
	assert.True(t, c1.Blend(c2).Equals(NewColor(.9, .2, .04)))
}
