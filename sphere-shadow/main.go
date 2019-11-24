package main

import (
	"fmt"

	"github.com/jefflinse/go-ray-tracer/rt"
)

func main() {
	rayOrigin := rt.NewPoint(0, 0, -5)
	wallZ := 5.0
	wallSize := 7.0
	canvasPixels := 100
	pixelSize := wallSize / float64(canvasPixels)
	halfWall := wallSize / 2

	canvas := rt.NewCanvas(canvasPixels, canvasPixels)
	color := rt.NewColor(1, 0, 0)
	shape := rt.NewSphere()

	for y := 0; y < canvasPixels; y++ {
		worldY := halfWall - pixelSize*float64(y)
		for x := 0; x < canvasPixels; x++ {
			worldX := -halfWall + pixelSize*float64(x)
			position := rt.NewPoint(worldX, worldY, wallZ)
			ray := rt.NewRay(rayOrigin, position.Subtract(rayOrigin).Normalize())
			intersections := shape.Intersect(ray)
			if intersections.Hit() != nil {
				canvas.WritePixel(x, y, color)
			}
		}
	}

	fmt.Print(canvas.ToPPM())
}
