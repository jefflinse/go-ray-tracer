package rt

import (
	"math"
)

// A Camera can be moved around and transformed to produce scenes.
type Camera struct {
	HSize      int
	VSize      int
	FOV        float64
	HalfWidth  float64
	HalfHeight float64
	PixelSize  float64
	Transform  Transformation
}

// NewCamera creates a new Camera
func NewCamera(hSize int, vSize int, fov float64) *Camera {
	camera := &Camera{
		HSize:     hSize,
		VSize:     vSize,
		FOV:       fov,
		Transform: NewTransform(),
	}

	halfView := math.Tan(camera.FOV / 2)
	aspectRatio := float64(camera.HSize) / float64(camera.VSize)
	if aspectRatio >= 1 {
		camera.HalfWidth = halfView
		camera.HalfHeight = halfView / aspectRatio
	} else {
		camera.HalfWidth = halfView * aspectRatio
		camera.HalfHeight = halfView
	}

	camera.PixelSize = (camera.HalfWidth * 2) / float64(camera.HSize)

	return camera
}

// RayForPixel returns a Ray that starts at the camera and passes through the pixel at x, y on the canvas.
func (c *Camera) RayForPixel(x int, y int) *Ray {
	px := float64(x)
	py := float64(y)
	xOffset := (px + .5) * c.PixelSize
	yOffset := (py + .5) * c.PixelSize
	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset
	pixel := c.Transform.Inverse().ApplyTo(NewPoint(worldX, worldY, -1))
	origin := c.Transform.Inverse().ApplyTo(Origin())
	direction := pixel.Subtract(origin).Normalize()
	return NewRay(origin, direction)
}

// Render renders the specified world.
func (c *Camera) Render(world *World) *Canvas {
	image := NewCanvas(c.HSize, c.VSize)
	for y := 0; y < c.VSize; y++ {
		for x := 0; x < c.HSize; x++ {
			ray := c.RayForPixel(x, y)
			color := world.ColorAt(ray)
			image.WritePixel(x, y, color)
		}
	}

	return image
}
