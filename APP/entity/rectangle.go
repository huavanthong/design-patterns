// rectangle.go
package entity

import (
	"fmt"
	"time"

	"github.com/huavanthong/design-patterns/APP/common"
)

type Rectangle struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	Dimensions common.Dimensions `json:"dimensions"`
	Color      common.Color      `json:"color"`
	Position   common.Position   `json:"position"`
	BorderSize int               `json:"border_size"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
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
