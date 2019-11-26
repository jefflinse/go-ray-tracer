package rt

import (
	"sort"
)

// An Intersection represents the intersection of a ray and an object.
type Intersection struct {
	T      float64
	Object Shape
}

// NewIntersection creates a new Intersection.
func NewIntersection(t float64, obj Shape) *Intersection {
	return &Intersection{t, obj}
}

// PrepareComputations precomputes intersection information.
func (i *Intersection) PrepareComputations(ray *Ray) *IntersectionInfo {
	info := &IntersectionInfo{
		Object: i.Object,
		T:      i.T,
	}

	info.Point = ray.Position(info.T)
	info.EyeV = ray.Direction.Negate()
	info.NormalV = i.Object.NormalAt(info.Point)

	if info.NormalV.Dot(info.EyeV) < 0 {
		info.Inside = true
		info.NormalV = info.NormalV.Negate()
	}

	info.OverPoint = info.Point.Add(info.NormalV.Multiply(EPSILON))

	return info
}

// An IntersectionInfo is a set of precomputed intersection information.
type IntersectionInfo struct {
	Object    Shape
	T         float64
	Point     Tuple
	OverPoint Tuple
	EyeV      Tuple
	NormalV   Tuple
	Inside    bool
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
	for _, x := range s {
		if x.T >= 0 {
			return x
		}
	}

	return nil
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
