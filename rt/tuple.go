package rt

import (
	"fmt"
	"math"
)

// A Tuple contains three dimensions and is either a Point or a Vector.
type Tuple struct {
	X float64
	Y float64
	Z float64
	T float64
}

// NewPoint creates a new Point.
func NewPoint(x float64, y float64, z float64) *Tuple {
	return &Tuple{x, y, z, 1}
}

// NewVector creates a new Vector.
func NewVector(x float64, y float64, z float64) *Tuple {
	return &Tuple{x, y, z, 0}
}

// Equals returns true if this tuple is equal to the other one.
func (t *Tuple) Equals(other *Tuple) bool {
	return eq(t.X, other.X) &&
		eq(t.Y, other.Y) &&
		eq(t.Z, other.Z) &&
		eq(t.T, other.T)
}

// Add creates a new Tuple by adding the values of this tuple with another one.
func (t *Tuple) Add(other *Tuple) *Tuple {
	return &Tuple{
		t.X + other.X,
		t.Y + other.Y,
		t.Z + other.Z,
		t.T + other.T,
	}
}

// Subtract creates a new Tuple by subtracting the values of another tuple from this one.
func (t *Tuple) Subtract(other *Tuple) *Tuple {
	return &Tuple{
		t.X - other.X,
		t.Y - other.Y,
		t.Z - other.Z,
		t.T - other.T,
	}
}

// Negate creates a new Tuple by negating the values of this tuple.
func (t *Tuple) Negate() *Tuple {
	return &Tuple{-t.X, -t.Y, -t.Z, -t.T}
}

// Multiply creates a new Tuple by multiplying this tuple by a scalar.
func (t *Tuple) Multiply(s float64) *Tuple {
	return &Tuple{t.X * s, t.Y * s, t.Z * s, t.T * s}
}

// Divide creates a new Tuple by dividing this tuple by a scalar.
func (t *Tuple) Divide(s float64) *Tuple {
	return t.Multiply(1 / s)
}

// Magnitude returns the magnitude of a vector.
func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2))
}

// Normalize creates a new Tuple by dividing each of this tuple's values by its magnitude.
func (t *Tuple) Normalize() *Tuple {
	mag := t.Magnitude()
	return &Tuple{t.X / mag, t.Y / mag, t.Z / mag, t.T}
}

// Dot returns the dot product of this vector and another.
func (t *Tuple) Dot(other *Tuple) float64 {
	return (t.X * other.X) + (t.Y * other.Y) + (t.Z * other.Z) + (t.T * other.T)
}

// Cross returns the cross product of this vector and another.
func (t *Tuple) Cross(other *Tuple) *Tuple {
	return &Tuple{
		(t.Y * other.Z) - (t.Z * other.Y),
		(t.Z * other.X) - (t.X * other.Z),
		(t.X * other.Y) - (t.Y * other.X),
		t.T,
	}
}

// String returns a string representation of a Tuple.
func (t *Tuple) String() string {
	return fmt.Sprintf("(%f,%f,%f)[%f]", t.X, t.Y, t.Z, t.T)
}
