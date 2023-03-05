// circle.go
package entity

import (
	"fmt"
	"math"

	"github.com/huavanthong/design-patterns/APP/common"
)

type Circle struct {
	radius     float64
	position   common.Position
	color      common.Color
	borderSize int
}

// Constructor
func NewCircle(radius float64, position common.Position, color common.Color, borderSize int) *Circle {
	return &Circle{radius: radius, position: position, color: color, borderSize: borderSize}
}

// Getter methods
func (c *Circle) Radius() float64 {
	return c.radius
}

func (c *Circle) Position() common.Position {
	return c.position
}
func (c *Circle) Color() common.Color {
	return c.color
}
func (c *Circle) BorderSize() int {
	return c.borderSize
}

func (c *Circle) SetRadius(radius float64) {
	c.radius = radius
}

func (c *Circle) SetPosition(position common.Position) {
	c.position = position
}

func (c *Circle) SetColor(color common.Color) {
	c.color = color
}

func (c *Circle) SetBorderSize(borderSize int) {
	c.borderSize = borderSize
}

// utilities
func (c *Circle) GetArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing a circle with radius %v, position %v, color %v, and border size %v\n", c.radius, c.position, c.color, c.borderSize)
}
