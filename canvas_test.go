package rt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCanvas(t *testing.T) {
	c := NewCanvas(10, 20)
	assert.Equal(t, 10, c.Width())
	assert.Equal(t, 20, c.Height())
}

func TestCanvas_WritePixel_PixelAt(t *testing.T) {
	c := NewCanvas(10, 20)
	red := NewColor(1, 0, 0)
	c.WritePixel(2, 3, red)
	assert.Equal(t, red, c.PixelAt(2, 3))
}

func TestCanvas_ToPPM(t *testing.T) {
	c := NewCanvas(5, 3)

	c.WritePixel(0, 0, NewColor(1.5, 0, 0))
	c.WritePixel(2, 1, NewColor(0, .5, 0))
	c.WritePixel(4, 2, NewColor(-.5, 0, 1))

	expected := "P3\n5 3\n255\n" +
		"255 0 0 0 0 0 0 0 0 0 0 0 0 0 0 " +
		"0 0 0 0 0 0 0 128 0 0 0 0 \n0 0 0 " +
		"0 0 0 0 0 0 0 0 0 0 0 0 0 0 255 \n"

	assert.Equal(t, expected, c.ToPPM())

	c = NewCanvas(10, 2)
	for j := range c.pixels {
		for i := range c.pixels[j] {
			c.WritePixel(i, j, NewColor(1, .8, .6))
		}
	}

	expected = "P3\n10 2\n255\n" +
		"255 204 153 255 204 153 255 204 153 255 204 153 \n" +
		"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 \n" +
		"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 \n" +
		"255 204 153 255 204 153 255 204 153 255 204 153 \n"

	assert.Equal(t, expected, c.ToPPM())
}
