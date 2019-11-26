package rt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIntersection(t *testing.T) {
	s := NewSphere()
	i := NewIntersection(3.5, s)
	assert.Equal(t, 3.5, i.T)
	assert.Equal(t, s, i.Object)
}

func TestIntersection_PrepareComputatations(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewSphere()
	i := NewIntersection(4, s)
	info := i.PrepareComputations(r)
	assert.Equal(t, i.T, info.T)
	assert.Equal(t, i.Object, info.Object)
	assert.Equal(t, NewPoint(0, 0, -1), info.Point)
	assert.Equal(t, NewVector(0, 0, -1), info.EyeV)
	assert.Equal(t, NewVector(0, 0, -1), info.NormalV)

	// the hit, when an intersection occurs on the outside of an object
	r = NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s = NewSphere()
	i = NewIntersection(4, s)
	info = i.PrepareComputations(r)
	assert.Equal(t, false, info.Inside)

	// the hit, when an intersection occurs on the inside of an object
	r = NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	s = NewSphere()
	i = NewIntersection(1, s)
	info = i.PrepareComputations(r)
	assert.Equal(t, NewPoint(0, 0, 1), info.Point)
	assert.Equal(t, NewVector(0, 0, -1), info.EyeV)
	assert.Equal(t, NewVector(0, 0, -1), info.NormalV)
	assert.Equal(t, true, info.Inside)
}

func TestNewIntersectionSet(t *testing.T) {
	xs := NewIntersectionSet()
	assert.Len(t, xs, 0)

	s := NewSphere()
	i1 := NewIntersection(5, s)
	i2 := NewIntersection(1, s)
	i3 := NewIntersection(-3, s)
	i4 := NewIntersection(4, s)
	xs = NewIntersectionSet(i1, i2, i3, i4)
	assert.Len(t, xs, 4)
	assert.Equal(t, i3, xs[0])
	assert.Equal(t, i2, xs[1])
	assert.Equal(t, i4, xs[2])
	assert.Equal(t, i1, xs[3])
}

func TestIntersectionSet_Hit(t *testing.T) {
	s := NewSphere()

	// all intersections have positive T
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)
	xs := NewIntersectionSet(i1, i2)
	assert.Equal(t, i1, xs.Hit())

	// some intersections have negative T
	i1 = NewIntersection(-1, s)
	i2 = NewIntersection(1, s)
	xs = NewIntersectionSet(i1, i2)
	assert.Equal(t, i2, xs.Hit())

	// all intersections have negative T
	i1 = NewIntersection(-2, s)
	i2 = NewIntersection(-1, s)
	xs = NewIntersectionSet(i1, i2)
	assert.Nil(t, xs.Hit())

	// hit is always the intersection with the lowest nonnegative t value
	i1 = NewIntersection(5, s)
	i2 = NewIntersection(7, s)
	i3 := NewIntersection(-3, s)
	i4 := NewIntersection(2, s)
	xs = NewIntersectionSet(i1, i2, i3, i4)
	assert.Equal(t, i4, xs.Hit())
}

func TestIntersectionSet_Join(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(5, s)
	i2 := NewIntersection(1, s)
	i3 := NewIntersection(-3, s)
	i4 := NewIntersection(4, s)

	xs1 := NewIntersectionSet(i1, i2)
	xs2 := NewIntersectionSet(i3, i4)
	xs3 := xs1.Join(xs2)
	assert.Len(t, xs3, 4)
	assert.Equal(t, i3, xs3[0])
	assert.Equal(t, i2, xs3[1])
	assert.Equal(t, i4, xs3[2])
	assert.Equal(t, i1, xs3[3])
}
