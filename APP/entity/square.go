// square.go
package entity

import (
	"fmt"
	"time"

	"github.com/huavanthong/design-patterns/APP/common"
)

type Square struct {
	ID         string          `json:"id"`
	ObjectName string          `json:"object_name"`
	OwnerName  string          `json:"owner_name"`
	Size       float64         `json:"size"`
	Color      common.Color    `json:"color"`
	Position   common.Position `json:"position"`
	BorderSize int             `json:"border_size"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}

// utilities
func (s Square) GetArea() float64 {
	return s.Size * s.Size
}

func (s Square) GetPerimeter() float64 {
	return 4 * (s.Size)
}

func (s Square) Draw() string {
	return fmt.Sprintf("Drawing Square with Size %v at position (%v, %v)", s.Size, s.Position.X, s.Position.Y)
}
