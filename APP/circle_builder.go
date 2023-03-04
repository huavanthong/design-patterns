package main

type CircleBuilder struct {
	shape *Circle
}

func NewCircleBuilder() *CircleBuilder {
	return &CircleBuilder{shape: &Circle{}}
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
