// File shape_builder.go
package main

type iShapeBuilder interface {
	SetDimensions(dimensions Dimensions)
	SetPosition(position Position)
	SetColor(color Color)
	Build() Shape
}

type Dimensions struct {
	Width  float64
	Height float64
	Radius float64
}

type Position struct {
	X float64
	Y float64
}

type Color struct {
	R uint8
	G uint8
	B uint8
}

func getShapeBuilder(builderType string) iShapeBuilder {
	if builderType == "Rectangle" {
		return &RectangleBuilder{}
	}
	if builderType == "Circle" {
		return &CircleBuilder{}
	}
	return nil
}
