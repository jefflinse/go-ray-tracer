package rt

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewShapeProps(t *testing.T) {
	sp := NewShapeProps()

	// has a default material
	assert.Equal(t, NewMaterial(), sp.Material)
	assert.Equal(t, NewMaterial(), sp.GetMaterial())

	// has a default transform
	assert.Equal(t, NewTransform(), sp.Transform)
}

func TestShapeProps_GetMaterial(t *testing.T) {
	sp := NewShapeProps()
	assert.Equal(t, NewMaterial(), sp.GetMaterial())
}

func TestShapeProps_intersect(t *testing.T) {
	// with default transform
	sp := NewShapeProps()
	r1 := NewRay(NewPoint(1, -2, 3), NewVector(4, 5, -6))
	sp.intersect(r1, func(localRay *Ray) IntersectionSet {
		assert.Equal(t, NewRay(NewPoint(1, -2, 3), NewVector(4, 5, -6)), localRay)
		return nil
	})

	// with custom transform
	sp = NewShapeProps()
	sp.Transform = NewTransform().Scale(.5, .5, .5)
	r1 = NewRay(NewPoint(1, -2, 3), NewVector(4, 5, -6))
	sp.intersect(r1, func(localRay *Ray) IntersectionSet {
		assert.Equal(t, NewRay(NewPoint(2, -4, 6), NewVector(8, 10, -12)), localRay)
		return nil
	})
}

func TestShapeProps_normalAt(t *testing.T) {

	// applies translation
	s := NewShapeProps()
	s.Transform = NewTranslation(0, 1, 0)
	p := NewPoint(0, 1.70711, -.70711)
	s.normalAt(p, func(localPoint Tuple) Tuple {
		assert.True(t, localPoint.Equals(NewPoint(0, .70711, -.70711)))
		return NewVector(1, 2, 3)
	})

	// applies scaling
	s = NewShapeProps()
	s.Transform = NewScaling(1, .5, 1).CombineWith(NewRotationZ(math.Pi / 5))
	p = NewPoint(0, math.Sqrt2/2, -math.Sqrt2/2)
	s.normalAt(p, func(localPoint Tuple) Tuple {
		assert.True(t, localPoint.Equals(NewPoint(.83125, 1.14412, -.70711)))
		return NewVector(1, 2, 3)
	})
}
