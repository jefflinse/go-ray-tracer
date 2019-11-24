package rt

// A Ray represents a ray in the ray tracer.
type Ray struct {
	Origin    Tuple
	Direction Tuple
}

// NewRay creates a new Ray.
func NewRay(origin Tuple, direction Tuple) *Ray {
	return &Ray{origin, direction}
}

// Position returns the point on the ray at distance t.
func (r *Ray) Position(t float64) Tuple {
	return r.Origin.Add(r.Direction.Multiply(t))
}
