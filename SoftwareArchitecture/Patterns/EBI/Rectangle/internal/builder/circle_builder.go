package builder

import (
	"github.com/huavanthong/design-patterns/SoftwareArchitecture/Patterns/EBI/Rectangle/internal/entity"
)

type CircleBuilder struct {
	shape *entity.Circle
}

func NewCircleBuilder() *CircleBuilder {
	return &CircleBuilder{shape: &entity.Circle{}}
}

func (cb *CircleBuilder) SetDimensions(dimensions Dimensions) {
	cb.shape.radius = dimensions.width / 2 // assuming width and height are the same
}

func (cb *CircleBuilder) SetPosition(position Position) {
	cb.shape.position = position
}

func (cb *CircleBuilder) SetColor(color Color) {
	cb.shape.color = color
}

func (cb *CircleBuilder) Build() Shape {
	return cb.shape
}
