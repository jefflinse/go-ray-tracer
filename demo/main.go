package main

import (
	"fmt"
	"math"

	"github.com/jefflinse/go-ray-tracer"
)

func main() {
	floor := rt.NewPlane()
	floor.Material = rt.NewMaterial()
	// floor.Material.Color = rt.NewColor(1, .9, .9)
	innerStripe := rt.NewStripePattern(
		rt.NewSolidPattern(rt.NewColor(.5, .5, .5)),
		rt.NewSolidPattern(rt.NewColor(0, 0, 0)),
	)
	innerStripe.SetTransform(rt.NewTransform().RotateY(math.Pi/4).Scale(.5, .5, .5))
	floor.Material.Pattern = rt.NewBlendedPattern(
		rt.NewSolidPattern(rt.NewColor(1, 0, 0)),
		innerStripe,
	)
	floor.Material.Specular = 0

	middle := rt.NewSphere()
	middle.Transform = rt.NewTransform().Translate(-.5, 1, .5)
	middle.Material = rt.NewMaterial()
	middle.Material.Color = rt.NewColor(.1, 1, .5)
	middle.Material.Diffuse = .7
	middle.Material.Specular = .3

	right := rt.NewSphere()
	right.Transform = rt.NewTransform().Scale(.5, .5, .5).Translate(1.5, .5, -.5)
	right.Material = rt.NewMaterial()
	right.Material.Color = rt.NewColor(.5, 1, .1)
	right.Material.Diffuse = .7
	right.Material.Specular = .3

	left := rt.NewSphere()
	left.Transform = rt.NewTransform().Scale(.33, .33, .33).Translate(-1.5, .33, -.75)
	left.Material = rt.NewMaterial()
	left.Material.Color = rt.NewColor(1, .8, .1)
	left.Material.Diffuse = .7
	left.Material.Specular = .3

	world := rt.NewWorld()
	world.Light = rt.NewPointLight(rt.NewPoint(-10, 10, -10), rt.NewColor(1, 1, 1))
	world.AddObjects(floor, middle, right, left)

	camera := rt.NewCamera(250, 125, math.Pi/3)
	camera.Transform = rt.NewViewTransform(
		rt.NewPoint(0, 1.5, -5),
		rt.NewPoint(0, 1, 0),
		rt.NewVector(0, 1, 0),
	)

	canvas := camera.Render(world)

	fmt.Print(canvas.ToPPM())
}
