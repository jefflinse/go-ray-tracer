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
