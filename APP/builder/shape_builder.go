// File shape_builder.go
package main

import (
	"github.com/huavanthong/design-patterns/APP/common"
)

type iShapeBuilder interface {
	SetDimensions(dimensions common.Dimensions)
	SetPosition(position common.Position)
	SetColor(color common.Color)
	Build() Shape
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
