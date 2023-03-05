// File shape_builder.go
package builder

import (
	"github.com/huavanthong/design-patterns/APP/common"
	"github.com/huavanthong/design-patterns/APP/entity"
)

type iShapeBuilder interface {
	SetDimensions(dimensions common.Dimensions)
	SetPosition(position common.Position)
	SetColor(color common.Color)
	Build() entity.Shape
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