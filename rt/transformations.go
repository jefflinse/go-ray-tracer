package rt

import (
	"math"
)

// NewTransform returns an identity matrix as the basis for a transformation.
func NewTransform() Matrix {
	return NewIdentityMatrix()
}

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

// Translate adds a translation to a transformation matrix.
func (m Matrix) Translate(x float64, y float64, z float64) Matrix {
	return NewTranslation(x, y, z).Multiply(m)
}

// Scale adds a scaling to a transformation matrix.
func (m Matrix) Scale(x float64, y float64, z float64) Matrix {
	return NewScaling(x, y, z).Multiply(m)
}

// RotateX adds an X-rotation to a transformation matrix.
func (m Matrix) RotateX(radians float64) Matrix {
	return NewRotationX(radians).Multiply(m)
}

// RotateY adds an Y-rotation to a transformation matrix.
func (m Matrix) RotateY(radians float64) Matrix {
	return NewRotationY(radians).Multiply(m)
}

// RotateZ adds an X-rotation to a transformation matrix.
func (m Matrix) RotateZ(radians float64) Matrix {
	return NewRotationZ(radians).Multiply(m)
}

// Shear adds a shearing to a transformation matrix.
func (m Matrix) Shear(xy float64, xz float64, yx float64, yz float64, zx float64, zy float64) Matrix {
	return NewShearing(xy, xz, yx, yz, zx, zy).Multiply(m)
}
