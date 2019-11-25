package rt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWorld(t *testing.T) {
	w := NewWorld()
	assert.NotNil(t, w.Objects)
}

func TestNewDefaultWorld(t *testing.T) {
	w := NewDefaultWorld()
	l := NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1))
	s1 := NewSphere()
	s1.Material.Color = NewColor(.8, 1, .6)
	s1.Material.Diffuse = .7
	s1.Material.Specular = .2
	s2 := NewSphere()
	s2.Transform = NewScaling(.5, .5, .5)
	assert.Equal(t, l, w.Light)
	assert.Equal(t, s1, w.Objects[0])
	assert.Equal(t, s2, w.Objects[1])
}

func TestWorld_AddObjects(t *testing.T) {
	w := NewWorld()
	w.AddObjects(NewSphere())
	assert.Len(t, w.Objects, 1)
	w.AddObjects(NewSphere(), NewSphere(), NewSphere())
	assert.Len(t, w.Objects, 4)
}

func TestWorld_Intersect(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	xs := w.Intersect(r)
	assert.Len(t, xs, 4)
	assert.Equal(t, xs[0].T, 4.0)
	assert.Equal(t, xs[1].T, 4.5)
	assert.Equal(t, xs[2].T, 5.5)
	assert.Equal(t, xs[3].T, 6.0)
}

func TestWorld_ShadeHit(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := w.Objects[0]
	i := NewIntersection(4, s)
	info := i.PrepareComputations(r)
	c := w.ShadeHit(info)
	assert.True(t, c.Equals(NewColor(.38066, .47583, .2855)))

	// shade intersection from inside
	w = NewDefaultWorld()
	w.Light = NewPointLight(NewPoint(0, .25, 0), NewColor(1, 1, 1))
	r = NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	s = w.Objects[1]
	i = NewIntersection(.5, s)
	info = i.PrepareComputations(r)
	c = w.ShadeHit(info)
	assert.True(t, c.Equals(NewColor(.90498, .90498, .90498)))
}

func TestWorld_ColorAt(t *testing.T) {
	// ray misses
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 1, 0))
	c := w.ColorAt(r)
	assert.True(t, c.Equals(NewColor(0, 0, 0)))

	// ray hits
	w = NewDefaultWorld()
	r = NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	c = w.ColorAt(r)
	assert.True(t, c.Equals(NewColor(.38066, .47583, .2855)))

	// intersection behind the ray
	w = NewDefaultWorld()
	outer := w.Objects[0].(*Sphere)
	outer.Material.Ambient = 1
	inner := w.Objects[1].(*Sphere)
	inner.Material.Ambient = 1
	r = NewRay(NewPoint(0, 0, .75), NewVector(0, 0, -1))
	c = w.ColorAt(r)
	assert.True(t, c.Equals(inner.Material.Color))
}
