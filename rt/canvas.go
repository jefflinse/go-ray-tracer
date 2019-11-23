package rt

import (
	"fmt"
	"strings"
)

// A Canvas is a grid of pixels.
type Canvas struct {
	width  int
	height int
	pixels [][]*Color
}

// NewCanvas creates a new Canvas.
func NewCanvas(width int, height int) *Canvas {
	pixels := make([][]*Color, height)
	for i := range pixels {
		pixels[i] = make([]*Color, width)
	}

	return &Canvas{width, height, pixels}
}

// Width returns the width of the canvas in pixels.
func (c *Canvas) Width() int {
	return c.width
}

// Height returns the height of the canvas in pixels.
func (c *Canvas) Height() int {
	return c.height
}

// PixelAt returns the Color at the specified position.
func (c *Canvas) PixelAt(x int, y int) *Color {
	return c.pixels[y][x]
}

// WritePixel writes a Color value to the specified position.
func (c *Canvas) WritePixel(x int, y int, color *Color) {
	c.pixels[y][x] = color
}

// ToPPM produces a PPM-formatted string from this canvas.
func (c *Canvas) ToPPM() string {
	builder := strings.Builder{}

	// write headers
	builder.WriteString("P3\n")
	builder.WriteString(fmt.Sprintf("%d %d\n", c.width, c.height))
	builder.WriteString("255\n")

	// write pixels
	for _, row := range c.pixels {
		for _, color := range row {
			pixelValue := color.ToPPM()
			if (builder.Len()%70)+len(pixelValue) >= 70 {
				builder.WriteByte('\n')
			}
			builder.WriteString(fmt.Sprintf("%s ", color.ToPPM()))
		}
	}

	builder.WriteByte('\n')

	return builder.String()
}
