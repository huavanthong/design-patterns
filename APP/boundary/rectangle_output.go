package boundary

import (
	"fmt"

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

// RectangleBoundary defines the boundary context for Rectangle.
type RectangleBoundary struct {
	interactor RectangleInteractor
}

// NewRectangleBoundary returns a new instance of RectangleBoundary.
func NewRectangleBoundary(interactor RectangleInteractor) *RectangleBoundary {
	return &RectangleBoundary{
		interactor: interactor,
	}
}

// CalculateRectangleArea calculates the area of the rectangle.
func CalculateRectangleArea(rectangle *entity.Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

// CalculateRectanglePerimeter calculates the perimeter of the rectangle.
func CalculateRectanglePerimeter(rectangle *entity.Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// SuccessRectangle is the success output method.
func (rb *RectangleBoundary) SuccessRectangle(output SuccessRectangleOutput) {
	fmt.Printf("Rectangle: %+v\n", output.Rectangle)
	fmt.Printf("Area: %f\n", output.Area)
	fmt.Printf("Perimeter: %f\n", output.Perimeter)
}

// ErrorRectangle is the error output method.
func (rb *RectangleBoundary) ErrorRectangle(err error) {
	fmt.Printf("Error: %s\n", err)
}
