package rt

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
	return c.pixels[x][y]
}

// WritePixel writes a Color value to the specified position.
func (c *Canvas) WritePixel(x int, y int, color *Color) {
	c.pixels[x][y] = color
}
