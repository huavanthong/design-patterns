package builder

import (
	"github.com/huavanthong/design-patterns/SoftwareArchitecture/Patterns/EBI/Rectangle/internal/entity"
)

type RectangleBuilder struct {
	shape *entity.Rectangle
}

func NewRectangleBuilder() *RectangleBuilder {
	return &RectangleBuilder{shape: &entity.Rectangle{}}
}

func (rb *RectangleBuilder) SetDimensions(dimensions Dimensions) {
	rb.shape.width = dimensions.width
	rb.shape.height = dimensions.height
}

func (rb *RectangleBuilder) SetPosition(position Position) {
	rb.shape.position = position
}

func (rb *RectangleBuilder) SetColor(color Color) {
	rb.shape.color = color
}

func (rb *RectangleBuilder) Build() Shape {
	return rb.shape
}
