package main

import (
	"fmt"
	"math"

	"github.com/jefflinse/go-ray-tracer/rt"
)

type environment struct {
	Gravity *rt.Tuple
	Wind    *rt.Tuple
}

type projectile struct {
	Position *rt.Tuple
	Velocity *rt.Tuple
}

func main() {
	canvas := rt.NewCanvas(900, 550)
	env := &environment{rt.NewVector(0, -0.1, 0), rt.NewVector(-0.01, 0, 0)}
	p := &projectile{rt.NewPoint(0, 1, 0), rt.NewVector(1, 1.8, 0).Normalize().Multiply(11.25)}
	for p.Position.Y > 0 {
		x, y := getPixelCoordinates(canvas, p)
		if x > 0 && x < canvas.Width() && y > 0 && y < canvas.Height() {
			canvas.WritePixel(x, y, rt.NewColor(1, 0, 0))
		}
		p = tick(env, p)
	}

	fmt.Print(canvas.ToPPM())
}

func getPixelCoordinates(canvas *rt.Canvas, p *projectile) (int, int) {
	x := int(math.Floor(p.Position.X))
	y := canvas.Height() - int(math.Floor(p.Position.Y))
	return x, y
}

func tick(env *environment, p *projectile) *projectile {
	return &projectile{
		p.Position.Add(p.Velocity),
		p.Velocity.Add(env.Gravity).Add(env.Wind),
	}
}
