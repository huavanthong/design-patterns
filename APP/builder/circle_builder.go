package builder

import (
	"github.com/huavanthong/design-patterns/APP/common"
	"github.com/huavanthong/design-patterns/APP/entity"
)

type CircleBuilder struct {
	shape *entity.Circle
}

func NewCircleBuilder() *CircleBuilder {
	return &CircleBuilder{shape: &entity.Circle{}}
}

func (cb *CircleBuilder) SetDimensions(dimensions common.Dimensions) {
	cb.shape.SetRadius(dimensions.Width / 2) // assuming width and height are the same
}

func (cb *CircleBuilder) SetPosition(position common.Position) {
	cb.shape.SetPosition(position)
}

func (cb *CircleBuilder) SetColor(color common.Color) {
	cb.shape.SetColor(color)
}

func (cb *CircleBuilder) Build() entity.Shape {
	return cb.shape.Shape
}
