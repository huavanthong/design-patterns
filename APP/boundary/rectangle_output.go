package boundary

import (
	"github.com/huavanthong/design-patterns/APP/entity"
)

// RectangleOutput defines the output port for Rectangle.
type RectangleOutput interface {
	SuccessRectangle(output SuccessRectangleOutput)
	ErrorRectangle(err error)
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
