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

func (c *Circle) Draw() {
	fmt.Printf("Drawing a circle with radius %v, position %v, color %v, and border size %v\n", c.radius, c.position, c.color, c.borderSize)
}

func (c *Circle) SetDimensions(dimensions common.Dimensions) {
	c.radius = dimensions.Radius
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

func (c *Circle) GetArea() float64 {
	return math.Pi * c.radius * c.radius
}
