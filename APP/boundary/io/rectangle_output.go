package boundary_output

import (
	"time"

	"github.com/huavanthong/design-patterns/APP/entity"
)

// RectangleOutputBoundary defines the output port for Rectangle.
type IRectangleOutput interface {
	SuccessRectangle(output SuccessRectangleOutput)
	ErrorRectangle(err error)
}

type RectangleOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Length    float64   `json:"length"`
	Breadth   float64   `json:"breadth"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SuccessRectangleOutput defines the success output structure for Rectangle.
type SuccessRectangleOutput struct {
	Rectangle *entity.Rectangle
	Area      float64
	Perimeter float64
}

// ErrorRectangleOutput defines the error output structure for Rectangle.
type ErrorRectangleOutput struct {
	Error string
}
