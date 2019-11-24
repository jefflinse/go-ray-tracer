package rt

// An Intersectable is anything that can be intersected by a ray.
type Intersectable interface {
	GetMaterial() Material
	NormalAt(worldPoint Tuple) Tuple
}

// An Intersection represents the intersection of a ray and an object.
type Intersection struct {
	T      float64
	Object Intersectable
}

// NewIntersection creates a new Intersection.
func NewIntersection(t float64, obj Intersectable) *Intersection {
	return &Intersection{t, obj}
}

// An IntersectionSet is a collection of Intersections.
type IntersectionSet []*Intersection

// Hit returns the intersection with the smallest nonnegative t value.
func (s IntersectionSet) Hit() *Intersection {
	var hit *Intersection
	for _, x := range s {
		if (hit == nil && x.T > 0) || (hit != nil && x.T > 0 && x.T < hit.T) {
			hit = x
		}
	}

	return hit
}
