// circle.go
package main

import (
	"fmt"
	"math"
)

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
