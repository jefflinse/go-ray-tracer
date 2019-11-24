package rt

import (
	"fmt"
	"math"
)

// A Tuple contains three dimensions and is either a Point or a Vector.
type Tuple []float64

// NewTuple creates a new zero-filled Tuple.
func NewTuple() Tuple {
	return make(Tuple, 4)
}

// NewPoint creates a new Point.
func NewPoint(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

// NewVector creates a new Vector.
func NewVector(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

// Origin returns a new Point at the origin.
func Origin() Tuple {
	return NewPoint(0, 0, 0)
}

// X returns the X value of the tuple.
func (t Tuple) X() float64 {
	return t[0]
}

// Y returns the Y value of the tuple.
func (t Tuple) Y() float64 {
	return t[1]
}

// Z returns the Z value of the tuple.
func (t Tuple) Z() float64 {
	return t[2]
}

// W returns the W value of the tuple.
func (t Tuple) W() float64 {
	return t[3]
}

// Equals returns true if this tuple is equal to the other one.
func (t Tuple) Equals(other Tuple) bool {
	return eq(t[0], other[0]) &&
		eq(t[1], other[1]) &&
		eq(t[2], other[2]) &&
		eq(t[3], other[3])
}

// Add creates a new Tuple by adding the values of this tuple with another one.
func (t Tuple) Add(other Tuple) Tuple {
	return Tuple{
		t[0] + other[0],
		t[1] + other[1],
		t[2] + other[2],
		t[3] + other[3],
	}
}

// Subtract creates a new Tuple by subtracting the values of another tuple from this one.
func (t Tuple) Subtract(other Tuple) Tuple {
	return Tuple{
		t[0] - other[0],
		t[1] - other[1],
		t[2] - other[2],
		t[3] - other[3],
	}
}

// Negate creates a new Tuple by negating the values of this tuple.
func (t Tuple) Negate() Tuple {
	return Tuple{-t[0], -t[1], -t[2], -t[3]}
}

// Multiply creates a new Tuple by multiplying this tuple by a scalar.
func (t Tuple) Multiply(s float64) Tuple {
	return Tuple{t[0] * s, t[1] * s, t[2] * s, t[3] * s}
}

// Divide creates a new Tuple by dividing this tuple by a scalar.
func (t Tuple) Divide(s float64) Tuple {
	return t.Multiply(1 / s)
}

// Magnitude returns the magnitude of a vector.
func (t Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t[0], 2) + math.Pow(t[1], 2) + math.Pow(t[2], 2))
}

// Normalize creates a new Tuple by dividing each of this tuple's values by its magnitude.
func (t Tuple) Normalize() *Tuple {
	mag := t.Magnitude()
	return &Tuple{t[0] / mag, t[1] / mag, t[2] / mag, t[3]}
}

// Dot returns the dot product of this vector and another.
func (t Tuple) Dot(other Tuple) float64 {
	return (t[0] * other[0]) + (t[1] * other[1]) + (t[2] * other[2]) + (t[3] * other[3])
}

// Cross returns the cross product of this vector and another.
func (t Tuple) Cross(other Tuple) Tuple {
	return Tuple{
		(t[1] * other[2]) - (t[2] * other[1]),
		(t[2] * other[0]) - (t[0] * other[2]),
		(t[0] * other[1]) - (t[1] * other[0]),
		t[3],
	}
}

// String returns a string representation of a Tuple.
func (t Tuple) String() string {
	return fmt.Sprintf("(%f,%f,%f)[%f]", t[0], t[1], t[2], t[3])
}
