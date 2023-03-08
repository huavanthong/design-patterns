// circle.go
package entity

import (
	"fmt"
	"math"
	"time"

	"github.com/huavanthong/design-patterns/APP/common"
)

type Circle struct {
	ID         int               `json:"id"`
	Name       string            `json:"name"`
	Dimensions common.Dimensions `json:"dimensions"`
	Color      common.Color      `json:"color"`
	Position   common.Position   `json:"position"`
	BorderSize int               `json:"border_size"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
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
