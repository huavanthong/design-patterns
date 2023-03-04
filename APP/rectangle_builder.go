package main

type RectangleBuilder struct {
	shape *Rectangle
}

func NewRectangleBuilder() *RectangleBuilder {
	return &RectangleBuilder{shape: &Rectangle{}}
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
