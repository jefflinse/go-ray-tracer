package rt

import (
	"math"
)

// NewTranslation creates a new translation matrix.
func NewTranslation(x float64, y float64, z float64) Matrix {
	m := NewIdentityMatrix()
	m[0][3] = x
	m[1][3] = y
	m[2][3] = z
	return m
}

// NewScaling creates a new scaling matrix.
func NewScaling(x float64, y float64, z float64) Matrix {
	m := NewIdentityMatrix()
	m[0][0] = x
	m[1][1] = y
	m[2][2] = z
	return m
}

// NewRotationX creates a new rotation matrix to rotate a point about the X axis.
func NewRotationX(radians float64) Matrix {
	m := NewIdentityMatrix()
	m[1][1] = math.Cos(radians)
	m[1][2] = -math.Sin(radians)
	m[2][1] = math.Sin(radians)
	m[2][2] = math.Cos(radians)
	return m
}

// NewRotationY creates a new rotation matrix to rotate a point about the Y axis.
func NewRotationY(radians float64) Matrix {
	m := NewIdentityMatrix()
	m[0][0] = math.Cos(radians)
	m[0][2] = math.Sin(radians)
	m[2][0] = -math.Sin(radians)
	m[2][2] = math.Cos(radians)
	return m
}

// NewRotationZ creates a new rotation matrix to rotate a point about the Z axis.
func NewRotationZ(radians float64) Matrix {
	m := NewIdentityMatrix()
	m[0][0] = math.Cos(radians)
	m[0][1] = -math.Sin(radians)
	m[1][0] = math.Sin(radians)
	m[1][1] = math.Cos(radians)
	return m
}

// NewShearing creates a new shearing matrix.
func NewShearing(xy float64, xz float64, yx float64, yz float64, zx float64, zy float64) Matrix {
	m := NewIdentityMatrix()
	m[0][1] = xy
	m[0][2] = xz
	m[1][0] = yx
	m[1][2] = yz
	m[2][0] = zx
	m[2][1] = zy
	return m
}
