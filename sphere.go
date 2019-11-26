package rt

import (
	"math"
)

// A Sphere represents a sphere.
type Sphere struct {
	ShapeProps
}

// NewSphere creates a new Sphere.
func NewSphere() *Sphere {
	return &Sphere{
		ShapeProps: NewShapeProps(),
	}
}

// Intersect returns a set of points where a ray intersects the sphere.
func (s *Sphere) Intersect(ray *Ray) IntersectionSet {
	return s.intersect(ray, func(localRay *Ray) IntersectionSet {
		sphereToRay := localRay.Origin.Subtract(Origin())
		a := localRay.Direction.Dot(localRay.Direction)
		b := 2 * localRay.Direction.Dot(sphereToRay)
		c := sphereToRay.Dot(sphereToRay) - 1

		discriminant := math.Pow(b, 2) - 4*a*c
		if discriminant < 0 {
			return IntersectionSet{}
		}

		return NewIntersectionSet(
			NewIntersection((-b-math.Sqrt(discriminant))/(2*a), s),
			NewIntersection((-b+math.Sqrt(discriminant))/(2*a), s),
		)
	})
}

// NormalAt returns the normal vector from the sphere for a point p.
func (s *Sphere) NormalAt(point Tuple) Tuple {
	return s.normalAt(point, func(localPoint Tuple) Tuple {
		localNormal := localPoint.Subtract(Origin())
		return localNormal
	})
}
