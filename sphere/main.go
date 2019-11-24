package main

import (
	"fmt"

	"github.com/jefflinse/go-ray-tracer/rt"
)

func main() {
	rayOrigin := rt.NewPoint(0, 0, -5)
	wallZ := 5.0
	wallSize := 7.0
	canvasPixels := 200
	pixelSize := wallSize / float64(canvasPixels)
	halfWall := wallSize / 2

	canvas := rt.NewCanvas(canvasPixels, canvasPixels)
	sphere := rt.NewSphere()
	sphere.Material.Color = rt.NewColor(1, .2, 1)
	light := rt.NewPointLight(rt.NewPoint(-10, 10, -10), rt.NewColor(1, 1, 1))

	for y := 0; y < canvasPixels; y++ {
		worldY := halfWall - pixelSize*float64(y)
		for x := 0; x < canvasPixels; x++ {
			worldX := -halfWall + pixelSize*float64(x)
			position := rt.NewPoint(worldX, worldY, wallZ)
			ray := rt.NewRay(rayOrigin, position.Subtract(rayOrigin).Normalize())
			intersections := sphere.Intersect(ray)
			if hit := intersections.Hit(); hit != nil {
				point := ray.Position(hit.T)
				normal := hit.Object.NormalAt(point)
				eye := ray.Direction.Negate()
				color := hit.Object.GetMaterial().Lighting(light, point, eye, normal)
				canvas.WritePixel(x, y, color)
			}
		}
	}

	fmt.Print(canvas.ToPPM())
}
