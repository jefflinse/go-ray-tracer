package rt

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTuple(t *testing.T) {
	tuple := NewTuple()
	assert.Len(t, tuple, 4)
}

func TestNewPoint(t *testing.T) {
	p := NewPoint(4.3, -4.2, 3.1)
	assert.Equal(t, p.X(), 4.3)
	assert.Equal(t, p.Y(), -4.2)
	assert.Equal(t, p.Z(), 3.1)
	assert.Equal(t, p.W(), 1.0)
}

func TestNewVector(t *testing.T) {
	p := NewVector(4.3, -4.2, 3.1)
	assert.Equal(t, p.X(), 4.3)
	assert.Equal(t, p.Y(), -4.2)
	assert.Equal(t, p.Z(), 3.1)
	assert.Equal(t, p.W(), 0.0)
}

func TestOrigin(t *testing.T) {
	o := Origin()
	assert.Equal(t, NewPoint(0, 0, 0), o)
}

func TestTuple_Equals(t *testing.T) {
	t1 := Tuple{1, 2, 3, 0}
	t2 := Tuple{1, 2, 3, 0}
	assert.True(t, t1.Equals(t2))
}

func TestTuple_Add(t *testing.T) {
	// add a vector to a vector to produce a vector
	t1 := Tuple{3, -2, 5, 0}
	t2 := Tuple{-2, 3, 1, 0}
	t3 := t1.Add(t2)

	// add a vector to a point to get a point
	assert.True(t, t3.Equals(Tuple{1, 1, 6, 0}))
	t1 = Tuple{3, -2, 5, 0}
	t2 = Tuple{-2, 3, 1, 1}
	t3 = t1.Add(t2)

	// add a point to a vector to get a point
	assert.True(t, t3.Equals(Tuple{1, 1, 6, 1}))
	t1 = Tuple{3, -2, 5, 1}
	t2 = Tuple{-2, 3, 1, 0}
	t3 = t1.Add(t2)
	assert.True(t, t3.Equals(Tuple{1, 1, 6, 1}))
}

func TestTuple_Subtract(t *testing.T) {
	// subtract a point from a point to get a vector
	p1 := NewPoint(3, 2, 1)
	p2 := NewPoint(5, 6, 7)
	p3 := p1.Subtract(p2)
	assert.True(t, p3.Equals(Tuple{-2, -4, -6, 0}))

	// subtract a vector from a point to get a point
	p1 = NewPoint(3, 2, 1)
	v1 := NewVector(5, 6, 7)
	p3 = p1.Subtract(v1)
	assert.True(t, p3.Equals(Tuple{-2, -4, -6, 1}))

	// subtract a vector from a vector to get a vector
	p1 = NewVector(3, 2, 1)
	v1 = NewVector(5, 6, 7)
	p3 = p1.Subtract(v1)
	assert.True(t, p3.Equals(Tuple{-2, -4, -6, 0}))
}

func TestTuple_Negate(t *testing.T) {
	t1 := &Tuple{1, -2, 3, -4}
	t2 := t1.Negate()
	assert.True(t, t2.Equals(Tuple{-1, 2, -3, 4}))
}

func TestTuple_Multiply(t *testing.T) {
	// multiply by s > 1
	t1 := &Tuple{1, -2, 3, -4}
	t2 := t1.Multiply(3.5)
	assert.True(t, t2.Equals(Tuple{3.5, -7, 10.5, -14}))

	// multiply by s < 1
	t1 = &Tuple{1, -2, 3, -4}
	t2 = t1.Multiply(.5)
	assert.True(t, t2.Equals(Tuple{.5, -1, 1.5, -2}))
}

func TestTuple_Divide(t *testing.T) {
	t1 := &Tuple{1, -2, 3, -4}
	t2 := t1.Divide(2)
	assert.True(t, t2.Equals(Tuple{.5, -1, 1.5, -2}))
}

func TestTuple_Magnitude(t *testing.T) {
	v1 := NewVector(1, 0, 0)
	assert.Equal(t, 1.0, v1.Magnitude())
	v1 = NewVector(0, 1, 0)
	assert.Equal(t, 1.0, v1.Magnitude())
	v1 = NewVector(0, 0, 1)
	assert.Equal(t, 1.0, v1.Magnitude())
	v1 = NewVector(1, 2, 3)
	assert.Equal(t, math.Sqrt(14), v1.Magnitude())
	v1 = NewVector(-1, -2, -3)
	assert.Equal(t, math.Sqrt(14), v1.Magnitude())
}

func TestTuple_Normalize(t *testing.T) {
	v1 := NewVector(4, 0, 0)
	v2 := v1.Normalize()
	assert.True(t, v2.Equals(NewVector(1, 0, 0)))
	v1 = NewVector(1, 2, 3)
	v2 = v1.Normalize()
	assert.True(t, v2.Equals(NewVector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14))))
	assert.Equal(t, 1.0, v2.Magnitude())
}

func TestTuple_Dot(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(2, 3, 4)
	assert.Equal(t, 20.0, v1.Dot(v2))
}

func TestTuple_Cross(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(2, 3, 4)
	assert.True(t, v1.Cross(v2).Equals(NewVector(-1, 2, -1)))
	assert.True(t, v2.Cross(v1).Equals(NewVector(1, -2, 1)))
}

func TestTuple_String(t *testing.T) {
	t1 := &Tuple{1, 2, 3, 4}
	assert.Equal(t, "(1.000000,2.000000,3.000000)[4.000000]", t1.String())
}
