package rt

import (
	"fmt"
	"math"
)

// A Color is a tuple of red, green, blue.
type Color struct {
	*Tuple
}

// NewColor creates a new Color.
func NewColor(red float64, green float64, blue float64) *Color {
	return &Color{&Tuple{red, green, blue, 0}}
}

// Red returns the red value of this color.
func (c *Color) Red() float64 {
	return c.X
}

// Green returns the green value of this color.
func (c *Color) Green() float64 {
	return c.Y
}

// Blue returns the blue value of this color.
func (c *Color) Blue() float64 {
	return c.Z
}

// Equals returns true if this color is equal to the other one.
func (c *Color) Equals(other *Color) bool {
	return c.Tuple.Equals(other.Tuple)
}

// Add creates a new Color by adding the components of this color with another one.
func (c *Color) Add(other *Color) *Color {
	return &Color{c.Tuple.Add(other.Tuple)}
}

// Subtract creates a new Color by subtracting the components of another color from this one.
func (c *Color) Subtract(other *Color) *Color {
	return &Color{c.Tuple.Subtract(other.Tuple)}
}

// Multiply creates a new Color by multiplying this color's components by a scalar.
func (c *Color) Multiply(s float64) *Color {
	return &Color{c.Tuple.Multiply(s)}
}

// Blend creates a new Color by blending this color with another color using the Hadamard product.
func (c *Color) Blend(other *Color) *Color {
	return NewColor(c.X*other.X, c.Y*other.Y, c.Z*other.Z)
}

// ToPPM retuns the PPM-formatted string representation of a Color.
func (c *Color) ToPPM() string {
	if c == nil {
		return "0 0 0"
	}

	max := float64(255)
	return fmt.Sprintf(
		"%d %d %d",
		int(math.Ceil(clamp(c.X*max, 0, max))),
		int(math.Ceil(clamp(c.Y*max, 0, max))),
		int(math.Ceil(clamp(c.Z*max, 0, max))),
	)
}
