package rt

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPointLight(t *testing.T) {
	position := Origin()
	intensity := NewColor(1, 1, 1)
	l := NewPointLight(position, intensity)
	assert.Equal(t, position, l.Position)
	assert.Equal(t, intensity, l.Intensity)
}

func TestNewMaterial(t *testing.T) {
	m := NewMaterial()
	assert.Equal(t, NewColor(1, 1, 1), m.Color)
	assert.Equal(t, .1, m.Ambient)
	assert.Equal(t, .9, m.Diffuse)
	assert.Equal(t, .9, m.Specular)
	assert.Equal(t, 200.0, m.Shininess)
}

func TestMaterial_Lighting(t *testing.T) {
	// dummpy object to satisfy call to Lighting()
	s := NewSphere()

	// eye between the light and the surface
	m := NewMaterial()
	p := Origin()
	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
	result := m.Lighting(s, light, p, eyeV, normalV, false)
	assert.True(t, result.Equals(NewColor(1.9, 1.9, 1.9)))

	// eye between the light and the surface, eye offset 45 degrees
	m = NewMaterial()
	p = Origin()
	eyeV = NewVector(0, math.Sqrt2/2, -math.Sqrt2/2)
	normalV = NewVector(0, 0, -1)
	light = NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
	result = m.Lighting(s, light, p, eyeV, normalV, false)
	assert.True(t, result.Equals(NewColor(1, 1, 1)))

	// eye opposite surface, light offset 45 degrees
	m = NewMaterial()
	p = Origin()
	eyeV = NewVector(0, 0, -1)
	normalV = NewVector(0, 0, -1)
	light = NewPointLight(NewPoint(0, 10, -10), NewColor(1, 1, 1))
	result = m.Lighting(s, light, p, eyeV, normalV, false)
	assert.True(t, result.Equals(NewColor(.7364, .7364, .7364)))

	// eye in the path of the reflection vector
	m = NewMaterial()
	p = Origin()
	eyeV = NewVector(0, -math.Sqrt2/2, -math.Sqrt2/2)
	normalV = NewVector(0, 0, -1)
	light = NewPointLight(NewPoint(0, 10, -10), NewColor(1, 1, 1))
	result = m.Lighting(s, light, p, eyeV, normalV, false)
	assert.True(t, result.Equals(NewColor(1.6364, 1.6364, 1.6364)))

	// light behind the surface
	m = NewMaterial()
	p = Origin()
	eyeV = NewVector(0, 0, -1)
	normalV = NewVector(0, 0, -1)
	light = NewPointLight(NewPoint(0, 0, 10), NewColor(1, 1, 1))
	result = m.Lighting(s, light, p, eyeV, normalV, false)
	assert.True(t, result.Equals(NewColor(.1, .1, .1)))

	// lighting with surface in shadow
	m = NewMaterial()
	p = Origin()
	eyeV = NewVector(0, 0, -1)
	normalV = NewVector(0, 0, -1)
	light = NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
	result = m.Lighting(s, light, p, eyeV, normalV, true)
	assert.True(t, result.Equals(NewColor(.1, .1, .1)))

	// lighting with a pattern applied
	m = NewMaterial()
	m.Pattern = NewStripePattern(NewColor(1, 1, 1), NewColor(0, 0, 0))
	m.Ambient = 1
	m.Diffuse = 0
	m.Specular = 0
	eyeV = NewVector(0, 0, -1)
	normalV = NewVector(0, 0, -1)
	light = NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
	c1 := m.Lighting(s, light, NewPoint(.9, 0, 0), eyeV, normalV, false)
	c2 := m.Lighting(s, light, NewPoint(1.1, 0, 0), eyeV, normalV, false)
	assert.Equal(t, NewColor(1, 1, 1), c1)
	assert.Equal(t, NewColor(0, 0, 0), c2)
}
