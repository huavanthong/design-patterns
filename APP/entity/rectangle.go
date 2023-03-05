// rectangle.go
package main

import (
	"fmt"
	"math"

	"github.com/huavanthong/design-patterns/APP/common"
)

type Rectangle struct {
	width      float64
	height     float64
	position   common.Position
	color      common.Color
	borderSize int
}

func (r *Rectangle) Draw() {
	fmt.Printf("Drawing a rectangle with width %v, height %v, position %v, color %v, and border size %v\n", r.width, r.height, r.position, r.color, r.borderSize)
}

func (r *Rectangle) SetDimensions(dimensions common.Dimensions) {
	r.width = dimensions.Width
	r.height = dimensions.Height
}

func (r *Rectangle) SetPosition(position common.osition) {
	r.position = position
}

func (r *Rectangle) SetColor(color common.Color) {
	r.color = color
}

func (r *Rectangle) SetBorderSize(borderSize int) {
	r.borderSize = borderSize
}

func (r *Rectangle) GetArea() float64 {
	return r.width * r.height
}

type Circle struct {
	radius     float64
	position   Position
	color      Color
	borderSize int
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing a circle with radius %v, position %v, color %v, and border size %v\n", c.radius, c.position, c.color, c.borderSize)
}

func (c *Circle) SetDimensions(dimensions Dimensions) {
	c.radius = dimensions.Radius
}

func (c *Circle) SetPosition(position Position) {
	c.position = position
}

func (c *Circle) SetColor(color Color) {
	c.color = color
}

func (c *Circle) SetBorderSize(borderSize int) {
	c.borderSize = borderSize
}

func (c *Circle) GetArea() float64 {
	return math.Pi * c.radius * c.radius
}
