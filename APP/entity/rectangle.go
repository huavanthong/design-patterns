// rectangle.go
package entity

import (
	"fmt"

	"github.com/huavanthong/design-patterns/APP/common"
)

type Rectangle struct {
	Dimensions common.Dimensions
	Color      common.Color
	Position   common.Position
	BorderSize int
}

// utilities
func (r Rectangle) GetArea() float64 {
	return r.Dimensions.Width * r.Dimensions.Height
}

func (r Rectangle) GetPerimeter() float64 {
	return 2 * (r.Dimensions.Width + r.Dimensions.Height)
}

func (r Rectangle) Draw() string {
	return fmt.Sprintf("Drawing Rectangle with width %v, height %v at position (%v, %v)", r.Dimensions.Width, r.Dimensions.Height, r.Position.X, r.Position.Y)
}
