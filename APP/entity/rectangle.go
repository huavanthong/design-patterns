// rectangle.go
package entity

import (
	"fmt"

	"github.com/huavanthong/design-patterns/APP/common"
)

type Rectangle struct {
	Shape
	width  float64
	height float64
}

// Constructor
func NewRectangle(width float64, height float64, position common.Position, color common.Color, borderSize int) *Rectangle {
	return &Rectangle{Shape: Shape{position: position, color: color, borderSize: borderSize}, width: width, height: height}
}

// Getter methods

func (r *Rectangle) Width() float64 {
	return r.width
}

func (r *Rectangle) Height() float64 {
	return r.height
}

func (r *Rectangle) Position() common.Position {
	return r.Shape.position
}

func (r *Rectangle) Color() common.Color {
	return r.Shape.color
}

func (r *Rectangle) BorderSize() int {
	return r.Shape.borderSize
}

// Setter methods

func (r *Rectangle) SetWidth(width float64) {
	r.width = width
}

func (r *Rectangle) SetHeight(height float64) {
	r.height = height
}

func (r *Rectangle) SetPosition(position common.Position) {
	r.Shape.position = position
}

func (r *Rectangle) SetColor(color common.Color) {
	r.Shape.color = color
}

func (r *Rectangle) SetBorderSize(borderSize int) {
	r.Shape.borderSize = borderSize
}

// utilities
func (r *Rectangle) GetArea() float64 {
	return r.width * r.height
}

func (r *Rectangle) GetPerimeter() float64 {
	return (r.width + r.height) * 2
}

func (r *Rectangle) Draw() {
	fmt.Printf("Drawing a rectangle with width %v, height %v, position %v, color %v, and border size %v\n", r.width, r.height, r.Shape.position, r.Shape.color, r.Shape.borderSize)
}
