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

// Transform applies the specified transformation matrix and returns a new ray.
func (r *Ray) Transform(m Matrix) *Ray {
	return NewRay(m.MultiplyTuple(r.Origin), m.MultiplyTuple(r.Direction))
}
