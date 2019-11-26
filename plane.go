package rt

import (
	"math"
)

// A Plane is an infinite 2D plane.
type Plane struct {
	ShapeProps
}

// NewPlane creates a new Plane.
func NewPlane() *Plane {
	return &Plane{
		ShapeProps: NewShapeProps(),
	}
}

// Intersect returns a set of points where a ray intersects the plane.
func (p *Plane) Intersect(ray *Ray) IntersectionSet {
	return p.intersect(ray, func(localRay *Ray) IntersectionSet {
		if math.Abs(ray.Direction.Y()) < EPSILON {
			return NewIntersectionSet()
		}

		t := -ray.Origin.Y() / ray.Direction.Y()
		return NewIntersectionSet(
			NewIntersection(t, p),
		)
	})
}

// NormalAt returns the normal vector from the plane for a point p.
func (p *Plane) NormalAt(point Tuple) Tuple {
	return p.normalAt(point, func(localPoint Tuple) Tuple {
		return NewVector(0, 1, 0)
	})
}
