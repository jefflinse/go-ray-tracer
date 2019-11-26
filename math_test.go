package rt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClamp(t *testing.T) {
	assert.Equal(t, float64(0), clamp(0, 0, 10))
	assert.Equal(t, float64(5), clamp(5, 0, 10))
	assert.Equal(t, float64(10), clamp(10, 0, 10))
	assert.Equal(t, float64(0), clamp(-5, 0, 10))
	assert.Equal(t, float64(10), clamp(15, 0, 10))
}

func TestEq(t *testing.T) {
	assert.True(t, eq(1, 1))
	assert.True(t, eq(1234567, 1234567))
	assert.True(t, eq(.12345, .12345))
	assert.True(t, eq(.0012345, .0012345))
	assert.True(t, eq(.00000001, .00000005))
	assert.False(t, eq(1234567, 1234568))
	assert.False(t, eq(.0012345, .0013345))
}
