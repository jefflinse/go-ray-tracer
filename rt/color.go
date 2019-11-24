package rt

import (
	"fmt"
	"math"
)

// A Color is a tuple of red, green, blue.
type Color Tuple

// NewColor creates a new Color.
func NewColor(red float64, green float64, blue float64) Color {
	return Color{red, green, blue, 0}
}

// Red returns the red value of this color.
func (c Color) Red() float64 {
	return c[0]
}

// Green returns the green value of this color.
func (c Color) Green() float64 {
	return c[1]
}

// Blue returns the blue value of this color.
func (c Color) Blue() float64 {
	return c[2]
}

// Equals returns true if this color is equal to another one.
func (c Color) Equals(other Color) bool {
	return Tuple(c).Equals(Tuple(other))
}

// Add creates a new Color by adding the components of this color with another one.
func (c Color) Add(other Color) Color {
	return Color(Tuple(c).Add(Tuple(other)))
}

// Subtract creates a new Color by subtracting the values of another color from this one.
func (c Color) Subtract(other Color) Color {
	return Color(Tuple(c).Subtract(Tuple(other)))
}

// Multiply creates a new Color by multiplying each component of this color by a scalar.
func (c Color) Multiply(s float64) Color {
	return NewColor(c[0]*s, c[1]*s, c[2]*s)
}

// Blend creates a new Color by blending this color with another color using the Hadamard product.
func (c Color) Blend(other Color) Color {
	return NewColor(c[0]*other[0], c[1]*other[1], c[2]*other[2])
}

// ToPPM retuns the PPM-formatted string representation of a Color.
func (c Color) ToPPM() string {
	if c == nil {
		return "0 0 0"
	}

	max := float64(255)
	return fmt.Sprintf(
		"%d %d %d",
		int(math.Ceil(clamp(c[0]*max, 0, max))),
		int(math.Ceil(clamp(c[1]*max, 0, max))),
		int(math.Ceil(clamp(c[2]*max, 0, max))),
	)
}
