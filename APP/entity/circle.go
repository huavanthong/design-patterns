// circle.go
package entity

import (
	"fmt"
	"math"

	"github.com/huavanthong/design-patterns/APP/common"
)

type Circle struct {
	Dimensions common.Dimensions
	Color      common.Color
	Position   common.Position
	BorderSize int
}

// utilities
func (c Circle) GetArea() float64 {
	return math.Pi * c.Dimensions.Radius * c.Dimensions.Radius
}

func (c Circle) GetPerimeter() float64 {
	return 2 * math.Pi * c.Dimensions.Radius
}

func (c Circle) Draw() string {
	return fmt.Sprintf("Drawing Circle with radius %v at position (%v, %v)", c.Dimensions.Radius, c.Position.X, c.Position.Y)
}
