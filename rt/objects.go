package rt

import (
	"math"
)

// A Sphere represents a sphere.
type Sphere struct{}

// NewSphere creates a new Sphere.
func NewSphere() *Sphere {
	return &Sphere{}
}

// Intersect returns a set of points where a ray intersects the sphere.
func (s *Sphere) Intersect(ray *Ray) IntersectionSet {
	sphereToRay := ray.Origin.Subtract(Origin())
	a := ray.Direction.Dot(ray.Direction)
	b := 2 * ray.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := math.Pow(b, 2) - 4*a*c
	if discriminant < 0 {
		return IntersectionSet{}
	}

	return IntersectionSet{
		NewIntersection((-b-math.Sqrt(discriminant))/2*a, s),
		NewIntersection((-b+math.Sqrt(discriminant))/2*a, s),
	}
}
