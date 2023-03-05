package main

type RectangleBuilder struct {
	shape *Rectangle
}

func NewRectangleBuilder() *RectangleBuilder {
	return &RectangleBuilder{shape: &Rectangle{}}
}

func (rb *RectangleBuilder) SetDimensions(dimensions Dimensions) {
	rb.shape.width = dimensions.Width
	rb.shape.height = dimensions.Height
}

func (rb *RectangleBuilder) SetPosition(position Position) {
	rb.shape.position = position
}

func (rb *RectangleBuilder) SetColor(color Color) {
	rb.shape.color = color
}

func (rb *RectangleBuilder) Build() *Rectangle {
	return rb.shape
}
