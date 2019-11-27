package rt

// A Shape is anything that can be rendered.
type Shape interface {
	GetMaterial() *Material
	GetTransform() Transformation
	Intersect(r *Ray) IntersectionSet
	NormalAt(p Tuple) Tuple
}

// ShapeProps contains properties common to all shapes.
type ShapeProps struct {
	Material  *Material
	Transform Transformation
}

// NewShapeProps creates a new ShapeProps.
func NewShapeProps() ShapeProps {
	return ShapeProps{
		Material:  NewMaterial(),
		Transform: NewTransform(),
	}
}

// GetMaterial gets the material properties.
func (sp *ShapeProps) GetMaterial() *Material {
	return sp.Material
}

// GetTransform gets the shape's transformation.
func (sp *ShapeProps) GetTransform() Transformation {
	return sp.Transform
}

func (sp *ShapeProps) intersect(worldRay *Ray, localIntersectFn func(localRay *Ray) IntersectionSet) IntersectionSet {
	localRay := worldRay.Transform(sp.Transform.Inverse())
	return localIntersectFn(localRay)
}

func (sp *ShapeProps) normalAt(worldPoint Tuple, localNormalFn func(localPoint Tuple) Tuple) Tuple {
	localPoint := sp.Transform.Inverse().ApplyTo(worldPoint)
	localNormal := localNormalFn(localPoint)
	worldNormal := sp.Transform.Inverse().Transpose().ApplyTo(localNormal)
	worldNormal[3] = 0
	return worldNormal.Normalize()
}
