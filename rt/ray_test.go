package rt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRay(t *testing.T) {
	origin := NewPoint(1, 2, 3)
	direction := NewVector(4, 5, 6)
	r := NewRay(origin, direction)
	assert.True(t, r.Origin.Equals(origin))
	assert.True(t, r.Direction.Equals(direction))
}

func TestRay_Position(t *testing.T) {
	r := NewRay(NewPoint(2, 3, 4), NewVector(1, 0, 0))
	assert.True(t, r.Position(0).Equals(NewPoint(2, 3, 4)))
	assert.True(t, r.Position(1).Equals(NewPoint(3, 3, 4)))
	assert.True(t, r.Position(-1).Equals(NewPoint(1, 3, 4)))
	assert.True(t, r.Position(2.5).Equals(NewPoint(4.5, 3, 4)))
}

func TestRay_Transform(t *testing.T) {
	r1 := NewRay(NewPoint(1, 2, 3), NewVector(0, 1, 0))

	// translate a ray
	m := NewTranslation(3, 4, 5)
	r2 := r1.Transform(m)
	assert.Equal(t, NewPoint(4, 6, 8), r2.Origin)
	assert.Equal(t, NewVector(0, 1, 0), r2.Direction)

	// scale a ray
	m = NewScaling(2, 3, 4)
	r2 = r1.Transform(m)
	assert.Equal(t, NewPoint(2, 6, 12), r2.Origin)
	assert.Equal(t, NewVector(0, 3, 0), r2.Direction)
}
