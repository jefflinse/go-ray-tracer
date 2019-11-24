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

func TestNewIntersectionSet(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)
	xs := IntersectionSet{i1, i2}
	assert.Len(t, xs, 2)
	assert.Equal(t, i1, xs[0])
	assert.Equal(t, i2, xs[1])
}

func TestIntersectionSet_Hit(t *testing.T) {
	s := NewSphere()

	// all intersections have positive T
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)
	xs := IntersectionSet{i1, i2}
	assert.Equal(t, i1, xs.Hit())

	// some intersections have negative T
	i1 = NewIntersection(-1, s)
	i2 = NewIntersection(1, s)
	xs = IntersectionSet{i1, i2}
	assert.Equal(t, i2, xs.Hit())

	// all intersections have negative T
	i1 = NewIntersection(-2, s)
	i2 = NewIntersection(-1, s)
	xs = IntersectionSet{i1, i2}
	assert.Nil(t, xs.Hit())

	// hit is always the intersection with the lowest nonnegative t value
	i1 = NewIntersection(5, s)
	i2 = NewIntersection(7, s)
	i3 := NewIntersection(-3, s)
	i4 := NewIntersection(2, s)
	xs = IntersectionSet{i1, i2, i3, i4}
	assert.Equal(t, i4, xs.Hit())
}
