package rt

import (
	"math"
)

// A PointLight is a light source originating from a single point.
type PointLight struct {
	Position  Tuple
	Intensity Color
}

// NewPointLight creates a new PointLight.
func NewPointLight(position Tuple, intensity Color) *PointLight {
	return &PointLight{position, intensity}
}

// Material describes the attributes of a material.
type Material struct {
	Ambient   float64
	Diffuse   float64
	Color     Color
	Specular  float64
	Shininess float64
}

// NewMaterial creates a new Material.
func NewMaterial() Material {
	return Material{
		Color:     NewColor(1, 1, 1),
		Ambient:   .1,
		Diffuse:   .9,
		Specular:  .9,
		Shininess: 200.0,
	}
}

// Lighting returns the computed color of the lighting for the given parameters.
func (m Material) Lighting(light *PointLight, position Tuple, eyeV Tuple, normalV Tuple, inShadow bool) Color {
	effectiveColor := m.Color.Blend(light.Intensity)
	lightV := light.Position.Subtract(position).Normalize()
	ambient := effectiveColor.Multiply(m.Ambient)
	lightDotNormal := lightV.Dot(normalV)
	diffuse := NewColor(0, 0, 0)
	specular := NewColor(0, 0, 0)
	if lightDotNormal >= 0 {
		diffuse = effectiveColor.Multiply(m.Diffuse).Multiply(lightDotNormal)
		reflectV := lightV.Negate().Reflect(normalV)
		reflectDotEye := reflectV.Dot(eyeV)
		if reflectDotEye > 0 {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = light.Intensity.Multiply(m.Specular).Multiply(factor)
		}
	}

	if inShadow {
		return ambient
	}

	return ambient.Add(diffuse).Add(specular)
}
