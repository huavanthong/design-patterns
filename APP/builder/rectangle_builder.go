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
	rb.shape.width = dimensions.Width
	rb.shape.height = dimensions.Height
}

func (rb *RectangleBuilder) SetPosition(position common.Position) {
	rb.shape.position = position
}

func (rb *RectangleBuilder) SetColor(color common.Color) {
	rb.shape.color = color
}

func (rb *RectangleBuilder) Build() *entity.Rectangle {
	return rb.shape
}
