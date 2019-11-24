package rt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSphere(t *testing.T) {
	s := NewSphere()
	assert.NotNil(t, s)
}

func TestRayIntersectsSphere(t *testing.T) {
	s := NewSphere()

	// intersection at two points
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	xs := s.Intersect(r)
	assert.Len(t, xs, 2)
	assert.True(t, eq(xs[0].T, 4.0))
	assert.Equal(t, xs[0].Object, s)
	assert.True(t, eq(xs[1].T, 6.0))
	assert.Equal(t, xs[1].Object, s)

	// intersection at a tangent
	r = NewRay(NewPoint(0, 1, -5), NewVector(0, 0, 1))
	xs = s.Intersect(r)
	assert.Len(t, xs, 2)
	assert.True(t, eq(xs[0].T, 5.0))
	assert.Equal(t, xs[0].Object, s)
	assert.True(t, eq(xs[1].T, 5.0))
	assert.Equal(t, xs[1].Object, s)

	// ray originates inside the sphere
	r = NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	xs = s.Intersect(r)
	assert.Len(t, xs, 2)
	assert.True(t, eq(xs[0].T, -1.0))
	assert.Equal(t, xs[0].Object, s)
	assert.True(t, eq(xs[1].T, 1.0))
	assert.Equal(t, xs[1].Object, s)

	// sphere is behind the ray
	r = NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
	xs = s.Intersect(r)
	assert.Len(t, xs, 2)
	assert.True(t, eq(xs[0].T, -6.0))
	assert.Equal(t, xs[0].Object, s)
	assert.True(t, eq(xs[1].T, -4.0))
	assert.Equal(t, xs[1].Object, s)

	// no intersection
	r = NewRay(NewPoint(2, 2, 2), NewVector(0, 0, 1))
	xs = s.Intersect(r)
	assert.Len(t, xs, 0)
}
