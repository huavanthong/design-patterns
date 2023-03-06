package builder

import (
	"github.com/huavanthong/design-patterns/APP/common"
	"github.com/huavanthong/design-patterns/APP/entity"
)

type RectangleBuilder struct {
	shape *entity.Rectangle
}

func NewRectangleBuilder() *RectangleBuilder {
	return &RectangleBuilder{shape: &entity.Rectangle{}}
}

func (rb *RectangleBuilder) SetDimensions(dimensions common.Dimensions) {
	rb.shape.SetWidth(dimensions.Width)
	rb.shape.SetHeight(dimensions.Height)
}

func (rb *RectangleBuilder) SetPosition(position common.Position) {
	rb.shape.SetPosition(position)
}

func (rb *RectangleBuilder) SetColor(color common.Color) {
	rb.shape.SetColor(color)
}

func (rb *RectangleBuilder) Build() entity.Shape {
	return rb.shape.Shape
}
