package rt

import (
	"sort"
)

// An Intersectable is anything that can be intersected by a ray.
type Intersectable interface {
	GetMaterial() Material
	Intersect(ray *Ray) IntersectionSet
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

// NewIntersectionSet creates a new IntersectionSet.
func NewIntersectionSet(xs ...*Intersection) IntersectionSet {
	set := IntersectionSet{}
	if xs != nil {
		set = set.Join(IntersectionSet(xs))
	}

	return set
}

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

// Join returns a new IntersectionSet containing all of the elements of the original set and the other one.
func (s IntersectionSet) Join(xs IntersectionSet) IntersectionSet {
	combined := IntersectionSet{}
	combined = append(combined, s...)
	combined = append(combined, xs...)
	sort.Sort(combined)
	return combined
}

// Functions to satisfy sorting interface.
func (s IntersectionSet) Len() int           { return len(s) }
func (s IntersectionSet) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntersectionSet) Less(i, j int) bool { return s[i].T < s[j].T }
