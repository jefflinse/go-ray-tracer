package rt

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSphere(t *testing.T) {
	s := NewSphere()
	assert.Equal(t, NewMaterial(), s.Material)
	assert.Equal(t, NewIdentityMatrix(), s.Transform)
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

func TestRayIntersectsSphereWithTransformation(t *testing.T) {
	// intersecting a scaled sphere
	s := NewSphere()
	s.Transform = NewScaling(2, 2, 2)
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	xs := s.Intersect(r)
	assert.Len(t, xs, 2)
	assert.Equal(t, 3.0, xs[0].T)
	assert.Equal(t, 7.0, xs[1].T)

	// intersecting a translated sphere
	s = NewSphere()
	s.Transform = NewTranslation(5, 0, 0)
	r = NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	xs = s.Intersect(r)
	assert.Len(t, xs, 0)
}

func TestSphere_NormalAt(t *testing.T) {
	s := NewSphere()

	// on the X axis
	n := s.NormalAt(NewPoint(1, 0, 0))
	assert.Equal(t, NewVector(1, 0, 0), n)

	// on the X axis
	n = s.NormalAt(NewPoint(0, 1, 0))
	assert.Equal(t, NewVector(0, 1, 0), n)

	// on the X axis
	n = s.NormalAt(NewPoint(0, 0, 1))
	assert.Equal(t, NewVector(0, 0, 1), n)

	// nonaxial point
	val := math.Sqrt(3) / 3
	n = s.NormalAt(NewPoint(val, val, val))
	assert.True(t, n.Equals(NewVector(val, val, val)))
	assert.True(t, n.Equals(n.Normalize()))

	// on a traslated sphere
	s = NewSphere()
	s.Transform = NewTranslation(0, 1, 0)
	n = s.NormalAt(NewPoint(0, 1.70711, -.70711))
	assert.True(t, n.Equals(NewVector(0, .70711, -.70711)))

	// on a transformed sphere
	s = NewSphere()
	s.Transform = NewScaling(1, .5, 1).Multiply(NewRotationZ(math.Pi / 5))
	n = s.NormalAt(NewPoint(0, math.Sqrt2/2, -math.Sqrt2/2))
	assert.True(t, n.Equals(NewVector(0, .97014, -.24254)))
}
