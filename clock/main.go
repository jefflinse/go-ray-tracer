package main

import (
	"fmt"
	"math"

	"github.com/jefflinse/go-ray-tracer/rt"
)

const radiansPerHour = math.Pi / 6

func main() {
	canvas := rt.NewCanvas(500, 500)
	radius := 200
	for hour := 0.0; hour < 12; hour++ {
		// 12-o-clock position
		p := rt.NewPoint(0, 1, 0)

		// rotate around Z axis to position on clock
		t := rt.NewTransform().RotateZ(hour * radiansPerHour)
		p = t.MultiplyTuple(p)

		x, y := getPixelCoordinates(canvas, radius, p)
		canvas.WritePixel(x, y, rt.NewColor(1, 1, 1))
	}

	fmt.Print(canvas.ToPPM())
}

func getPixelCoordinates(canvas *rt.Canvas, radius int, p rt.Tuple) (int, int) {
	centerX := canvas.Width() / 2
	centerY := canvas.Height() / 2
	x := centerX + int(math.Floor(p.X()*float64(radius)))
	y := centerY - int(math.Floor(p.Y()*float64(radius)))
	return x, y
}
