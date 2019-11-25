package rt

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCamera(t *testing.T) {
	c := NewCamera(160, 120, math.Pi/2)
	assert.Equal(t, 160, c.HSize)
	assert.Equal(t, 120, c.VSize)
	assert.Equal(t, math.Pi/2, c.FOV)
	assert.Equal(t, NewTransform(), c.Transform)
}

func TestCamera_GetPixelSize(t *testing.T) {
	// horizontal canvas
	c := NewCamera(200, 125, math.Pi/2)
	assert.Equal(t, .01, c.PixelSize)

	// vertical canvas
	c = NewCamera(125, 200, math.Pi/2)
	assert.Equal(t, .01, c.PixelSize)
}

func TestCamera_RayForPixel(t *testing.T) {
	// ray through center of the canvas
	c := NewCamera(201, 101, math.Pi/2)
	r := c.RayForPixel(100, 50)
	assert.True(t, r.Origin.Equals(Origin()))
	assert.True(t, r.Direction.Equals(NewVector(0, 0, -1)))

	// ray through a corner of the canvas
	c = NewCamera(201, 101, math.Pi/2)
	r = c.RayForPixel(0, 0)
	assert.True(t, r.Origin.Equals(Origin()))
	assert.True(t, r.Direction.Equals(NewVector(.66519, .33259, -.66851)))

	// ray with transformed camera
	c = NewCamera(201, 101, math.Pi/2)
	c.Transform = NewRotationY(math.Pi / 4).CombineWith(NewTranslation(0, -2, 5))
	r = c.RayForPixel(100, 50)
	assert.True(t, r.Origin.Equals(NewPoint(0, 2, -5)))
	assert.True(t, r.Direction.Equals(NewVector(math.Sqrt2/2, 0, -math.Sqrt2/2)))
}

func TestCamera_Render(t *testing.T) {
	w := NewDefaultWorld()
	c := NewCamera(11, 11, math.Pi/2)
	from := NewPoint(0, 0, -5)
	to := Origin()
	up := NewVector(0, 1, 0)
	c.Transform = NewViewTransform(from, to, up)
	image := c.Render(w)
	assert.True(t, image.PixelAt(5, 5).Equals(NewColor(.38066, .47583, .2855)))
}
