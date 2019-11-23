package rt

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslation(t *testing.T) {
	tm := NewTranslation(5, -3, 2)
	p := NewPoint(-3, 4, 5)

	// translate a point
	assert.True(t, tm.MultiplyTuple(p).Equals(NewPoint(2, 1, 7)))

	// translate a point (inverse)
	itm := tm.Inverse()
	assert.True(t, itm.MultiplyTuple(p).Equals(NewPoint(-8, 7, 3)))

	// translation does not affect vectors
	v := NewVector(-3, 4, 5)
	assert.True(t, tm.MultiplyTuple(v).Equals(v))
}

func TestScaling(t *testing.T) {
	sm := NewScaling(2, 3, 4)
	p := NewPoint(-4, 6, 8)

	// scale a point
	assert.True(t, sm.MultiplyTuple(p).Equals(NewPoint(-8, 18, 32)))

	// scale a vector
	v := NewVector(-4, 6, 8)
	assert.True(t, sm.MultiplyTuple(v).Equals(NewVector(-8, 18, 32)))

	// scale a vector (inverse)
	ism := sm.Inverse()
	assert.True(t, ism.MultiplyTuple(v).Equals(NewVector(-2, 2, 2)))

	// reflection is scaling by negative value
	sm = NewScaling(-1, 1, 1)
	p = NewPoint(2, 3, 4)
	assert.True(t, sm.MultiplyTuple(p).Equals(NewPoint(-2, 3, 4)))
}

func TestRotationX(t *testing.T) {
	p := NewPoint(0, 1, 0)
	halfQuarter := NewRotationX(math.Pi / 4)
	fullQuarter := NewRotationX(math.Pi / 2)
	assert.True(t, halfQuarter.MultiplyTuple(p).Equals(NewPoint(0, math.Sqrt2/2, math.Sqrt2/2)))
	assert.True(t, fullQuarter.MultiplyTuple(p).Equals(NewPoint(0, 0, 1)))

	// inverse rotation
	assert.True(t, halfQuarter.Inverse().MultiplyTuple(p).Equals(NewPoint(0, math.Sqrt2/2, -math.Sqrt2/2)))
}

func TestRotationY(t *testing.T) {
	p := NewPoint(0, 0, 1)
	halfQuarter := NewRotationY(math.Pi / 4)
	fullQuarter := NewRotationY(math.Pi / 2)
	assert.True(t, halfQuarter.MultiplyTuple(p).Equals(NewPoint(math.Sqrt2/2, 0, math.Sqrt2/2)))
	assert.True(t, fullQuarter.MultiplyTuple(p).Equals(NewPoint(1, 0, 0)))

	// inverse rotation
	assert.True(t, halfQuarter.Inverse().MultiplyTuple(p).Equals(NewPoint(-math.Sqrt2/2, 0, math.Sqrt2/2)))
}

func TestRotationZ(t *testing.T) {
	p := NewPoint(0, 1, 0)
	halfQuarter := NewRotationZ(math.Pi / 4)
	fullQuarter := NewRotationZ(math.Pi / 2)
	assert.True(t, halfQuarter.MultiplyTuple(p).Equals(NewPoint(-math.Sqrt2/2, math.Sqrt2/2, 0)))
	assert.True(t, fullQuarter.MultiplyTuple(p).Equals(NewPoint(-1, 0, 0)))

	// inverse rotation
	assert.True(t, halfQuarter.Inverse().MultiplyTuple(p).Equals(NewPoint(math.Sqrt2/2, math.Sqrt2/2, 0)))
}

func TestShearing(t *testing.T) {
	p := NewPoint(2, 3, 4)

	// x in proportion to y
	m := NewShearing(1, 0, 0, 0, 0, 0)
	assert.True(t, m.MultiplyTuple(p).Equals(NewPoint(5, 3, 4)))

	// x in proportion to z
	m = NewShearing(0, 1, 0, 0, 0, 0)
	assert.True(t, m.MultiplyTuple(p).Equals(NewPoint(6, 3, 4)))

	// y in proportion to x
	m = NewShearing(0, 0, 1, 0, 0, 0)
	assert.True(t, m.MultiplyTuple(p).Equals(NewPoint(2, 5, 4)))

	// y in proportion to z
	m = NewShearing(0, 0, 0, 1, 0, 0)
	assert.True(t, m.MultiplyTuple(p).Equals(NewPoint(2, 7, 4)))

	// z in proportion to x
	m = NewShearing(0, 0, 0, 0, 1, 0)
	assert.True(t, m.MultiplyTuple(p).Equals(NewPoint(2, 3, 6)))

	// z in proportion to y
	m = NewShearing(0, 0, 0, 0, 0, 1)
	assert.True(t, m.MultiplyTuple(p).Equals(NewPoint(2, 3, 7)))
}

func TestSqeuentialTransformations(t *testing.T) {
	p := NewPoint(1, 0, 1)
	rotate := NewRotationX(math.Pi / 2)
	scale := NewScaling(5, 5, 5)
	translate := NewTranslation(10, 5, 7)

	p2 := rotate.MultiplyTuple(p)
	assert.True(t, p2.Equals(NewPoint(1, -1, 0)))
	p3 := scale.MultiplyTuple(p2)
	assert.True(t, p3.Equals(NewPoint(5, -5, 0)))
	p4 := translate.MultiplyTuple(p3)
	assert.True(t, p4.Equals(NewPoint(15, 0, 7)))

	// chained transformations are applied in reverse order
	chained := translate.Multiply(scale).Multiply(rotate)
	assert.True(t, chained.MultiplyTuple(p).Equals(NewPoint(15, 0, 7)))
}

func TestTransformationAPIs(t *testing.T) {
	assert.True(t, NewTransform().Translate(1, 1, 1).Equals(NewIdentityMatrix().Multiply(NewTranslation(1, 1, 1))))
	assert.True(t, NewTransform().Scale(1, 1, 1).Equals(NewIdentityMatrix().Multiply(NewScaling(1, 1, 1))))
	assert.True(t, NewTransform().RotateX(.5).Equals(NewIdentityMatrix().Multiply(NewRotationX(.5))))
	assert.True(t, NewTransform().RotateY(.5).Equals(NewIdentityMatrix().Multiply(NewRotationY(.5))))
	assert.True(t, NewTransform().RotateZ(.5).Equals(NewIdentityMatrix().Multiply(NewRotationZ(.5))))
	assert.True(t, NewTransform().Shear(1, 1, 1, 1, 1, 1).Equals(NewIdentityMatrix().Multiply(NewShearing(1, 1, 1, 1, 1, 1))))

	// chaining works as expected
	rotate := NewRotationX(math.Pi / 2)
	scale := NewScaling(5, 5, 5)
	translate := NewTranslation(10, 5, 7)
	t1 := translate.Multiply(scale).Multiply(rotate)
	t2 := NewTransform().RotateX(math.Pi/2).Scale(5, 5, 5).Translate(10, 5, 7)
	assert.True(t, t2.Equals(t1))
}
