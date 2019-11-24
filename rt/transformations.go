package rt

import (
	"math"
)

// A Transformation is a Matrix that can be applied to Tuples.
type Transformation Matrix

// NewTransform returns an identity matrix as the basis for a transformation.
func NewTransform() Transformation {
	return Transformation(NewIdentityMatrix())
}

// NewTranslation creates a new translation matrix.
func NewTranslation(x float64, y float64, z float64) Transformation {
	m := NewIdentityMatrix()
	m[0][3] = x
	m[1][3] = y
	m[2][3] = z
	return Transformation(m)
}

// NewScaling creates a new scaling matrix.
func NewScaling(x float64, y float64, z float64) Transformation {
	m := NewIdentityMatrix()
	m[0][0] = x
	m[1][1] = y
	m[2][2] = z
	return Transformation(m)
}

// NewRotationX creates a new rotation matrix to rotate a point about the X axis.
func NewRotationX(radians float64) Transformation {
	m := NewIdentityMatrix()
	m[1][1] = math.Cos(radians)
	m[1][2] = -math.Sin(radians)
	m[2][1] = math.Sin(radians)
	m[2][2] = math.Cos(radians)
	return Transformation(m)
}

// NewRotationY creates a new rotation matrix to rotate a point about the Y axis.
func NewRotationY(radians float64) Transformation {
	m := NewIdentityMatrix()
	m[0][0] = math.Cos(radians)
	m[0][2] = math.Sin(radians)
	m[2][0] = -math.Sin(radians)
	m[2][2] = math.Cos(radians)
	return Transformation(m)
}

// NewRotationZ creates a new rotation matrix to rotate a point about the Z axis.
func NewRotationZ(radians float64) Transformation {
	m := NewIdentityMatrix()
	m[0][0] = math.Cos(radians)
	m[0][1] = -math.Sin(radians)
	m[1][0] = math.Sin(radians)
	m[1][1] = math.Cos(radians)
	return Transformation(m)
}

// NewShearing creates a new shearing matrix.
func NewShearing(xy float64, xz float64, yx float64, yz float64, zx float64, zy float64) Transformation {
	m := NewIdentityMatrix()
	m[0][1] = xy
	m[0][2] = xz
	m[1][0] = yx
	m[1][2] = yz
	m[2][0] = zx
	m[2][1] = zy
	return Transformation(m)
}

// ApplyTo applies the transformation matrix to the specified tuple to produce a new tuple.
func (t Transformation) ApplyTo(p Tuple) Tuple {
	return Matrix(t).MultiplyTuple(p)
}

// CombineWith combines this transformation with another one.
func (t Transformation) CombineWith(other Transformation) Transformation {
	return Transformation(Matrix(t).Multiply(Matrix(other)))
}

// Equals returns true if this transformation is equal to another one.
func (t Transformation) Equals(other Transformation) bool {
	return Matrix(t).Equals(Matrix(other))
}

// Inverse returns the inverse of this transformation.
func (t Transformation) Inverse() Transformation {
	return Transformation(Matrix(t).Inverse())
}

// Transpose returns a transposed version of this transformation.
func (t Transformation) Transpose() Transformation {
	return Transformation(Matrix(t).Transpose())
}

// Translate chains the current transformation with a translation.
func (t Transformation) Translate(x float64, y float64, z float64) Transformation {
	return NewTranslation(x, y, z).CombineWith(t)
}

// Scale chains the current transformation with a scaling.
func (t Transformation) Scale(x float64, y float64, z float64) Transformation {
	return NewScaling(x, y, z).CombineWith(t)
}

// RotateX chains the current transformation with an X-rotation.
func (t Transformation) RotateX(radians float64) Transformation {
	return NewRotationX(radians).CombineWith(t)
}

// RotateY chains the current transformation with an Y-rotation.
func (t Transformation) RotateY(radians float64) Transformation {
	return NewRotationY(radians).CombineWith(t)
}

// RotateZ chains the current transformation with an Z-rotation.
func (t Transformation) RotateZ(radians float64) Transformation {
	return NewRotationZ(radians).CombineWith(t)
}

// Shear chains the current transformation with a shearing.
func (t Transformation) Shear(xy float64, xz float64, yx float64, yz float64, zx float64, zy float64) Transformation {
	return NewShearing(xy, xz, yx, yz, zx, zy).CombineWith(t)
}
