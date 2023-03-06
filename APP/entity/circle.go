// circle.go
package entity

import (
	"fmt"
	"math"

	"github.com/huavanthong/design-patterns/APP/common"
)

type Circle struct {
	Shape
	radius float64
}

// Constructor
func NewCircle(radius float64, position common.Position, color common.Color, borderSize int) *Circle {
	return &Circle{Shape: Shape{position: position, color: color, borderSize: borderSize}, radius: radius}
}

// Getter methods
func (c *Circle) Radius() float64 {
	return c.radius
}

func (c *Circle) Position() common.Position {
	return c.Shape.position
}
func (c *Circle) Color() common.Color {
	return c.Shape.color
}
func (c *Circle) BorderSize() int {
	return c.Shape.borderSize
}

func (c *Circle) SetRadius(radius float64) {
	c.radius = radius
}

func (c *Circle) SetPosition(position common.Position) {
	c.Shape.position = position
}

func (c *Circle) SetColor(color common.Color) {
	c.Shape.color = color
}

func (c *Circle) SetBorderSize(borderSize int) {
	c.Shape.borderSize = borderSize
}

// utilities
func (c *Circle) GetArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) GetPerimeter() float64 {
	return math.Pi * 2 * c.radius
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing a circle with radius %v, position %v, color %v, and border size %v\n", c.radius, c.Shape.position, c.Shape.color, c.Shape.borderSize)
}
