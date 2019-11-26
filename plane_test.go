package rt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlane(t *testing.T) {
	p := NewPlane()
	assert.NotNil(t, p)
}

func TestPlane_Intersect(t *testing.T) {
	// ray is parallel to the plane
	p := NewPlane()
	r := NewRay(NewPoint(0, 10, 0), NewVector(0, 0, 1))
	xs := p.Intersect(r)
	assert.Empty(t, xs)

	// ray is coplanar
	p = NewPlane()
	r = NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	xs = p.Intersect(r)
	assert.Empty(t, xs)

	// ray intersects the plane from above
	p = NewPlane()
	r = NewRay(NewPoint(0, 1, 0), NewVector(0, -1, 0))
	xs = p.Intersect(r)
	assert.Len(t, xs, 1)
	assert.Equal(t, 1.0, xs[0].T)
	assert.Equal(t, p, xs[0].Object)

	// ray intersects the plane from below
	p = NewPlane()
	r = NewRay(NewPoint(0, -1, 0), NewVector(0, 1, 0))
	xs = p.Intersect(r)
	assert.Len(t, xs, 1)
	assert.Equal(t, 1.0, xs[0].T)
	assert.Equal(t, p, xs[0].Object)
}

func TestPlane_NormalAt(t *testing.T) {
	p := NewPlane()
	assert.True(t, p.NormalAt(Origin()).Equals(NewVector(0, 1, 0)))
	assert.True(t, p.NormalAt(NewPoint(10, 0, -10)).Equals(NewVector(0, 1, 0)))
	assert.True(t, p.NormalAt(NewPoint(-5, 0, 150)).Equals(NewVector(0, 1, 0)))
}
