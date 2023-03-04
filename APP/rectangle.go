// rectangle.go
package entity

import (
	"fmt"
)

type Rectangle struct {
	width      float64
	height     float64
	position   Position
	color      Color
	borderSize int
}

func (r *Rectangle) Draw() {
	fmt.Printf("Drawing a rectangle with width %v, height %v, position %v, color %v, and border size %v\n", r.width, r.height, r.position, r.color, r.borderSize)
}

func (r *Rectangle) SetDimensions(dimensions Dimensions) {
	r.width = dimensions.Width
	r.height = dimensions.Height
}

func (r *Rectangle) SetPosition(position Position) {
	r.position = position
}

func (r *Rectangle) SetColor(color Color) {
	r.color = color
}

func (r *Rectangle) SetBorderSize(borderSize int) {
	r.borderSize = borderSize
}

func (r *Rectangle) GetArea() float64 {
	return r.width * r.height
}
