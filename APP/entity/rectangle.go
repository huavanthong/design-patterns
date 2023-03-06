// rectangle.go
package entity

import (
	"fmt"

	"github.com/huavanthong/design-patterns/APP/common"
)

type Rectangle struct {
	width      float64
	height     float64
	position   common.Position
	color      common.Color
	borderSize int
}

// Constructor
func NewRectangle(width float64, height float64, position common.Position, color common.Color, borderSize int) *Rectangle {
	return &Rectangle{width: width, height: height, position: position, color: color, borderSize: borderSize}
}

// Getter methods

func (r *Rectangle) Width() float64 {
	return r.width
}

func (r *Rectangle) Height() float64 {
	return r.height
}

func (r *Rectangle) Position() common.Position {
	return r.position
}

func (r *Rectangle) Color() common.Color {
	return r.color
}

func (r *Rectangle) BorderSize() int {
	return r.borderSize
}

// Setter methods

func (r *Rectangle) SetWidth(width float64) {
	r.width = width
}

func (r *Rectangle) SetHeight(height float64) {
	r.height = height
}

func (r *Rectangle) SetPosition(position common.Position) {
	r.position = position
}

func (r *Rectangle) SetColor(color common.Color) {
	r.color = color
}

func (r *Rectangle) SetBorderSize(borderSize int) {
	r.borderSize = borderSize
}

// utilities
func (r *Rectangle) GetArea() float64 {
	return r.width * r.height
}

func (r *Rectangle) GetPerimeter() float64 {
	return (r.width + r.height) * 2
}

func (r *Rectangle) Draw() {
	fmt.Printf("Drawing a rectangle with width %v, height %v, position %v, color %v, and border size %v\n", r.width, r.height, r.position, r.color, r.borderSize)
}
