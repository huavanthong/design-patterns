package boundary

import (
	"fmt"

	"github.com/huavanthong/design-patterns/APP/entity"
)

// RectangleBoundary defines the boundary context for Rectangle.
type RectangleBoundary struct {
	input  RectangleInput
	output RectangleOutput
}

// NewRectangleBoundary is a factory function
// that it returns a new instance of RectangleBoundary.
func NewRectangleBoundary(input RectangleInput, output RectangleOutput) *RectangleBoundary {
	return &RectangleBoundary{
		input:  input,
		output: output,
	}
}

// CreateRectangle creates a rectangle.
func (rb *RectangleBoundary) CreateRectangle(input CreateRectangleInput) (*entity.Rectangle, error) {
	rectangle, err := rb.input.CreateRectangle(input)
	if err != nil {
		rb.output.ErrorRectangle(err)
		return nil, err
	}

	area := CalculateRectangleArea(rectangle)
	perimeter := CalculateRectanglePerimeter(rectangle)

	rb.output.SuccessRectangle(SuccessRectangleOutput{
		Rectangle: rectangle,
		Area:      area,
		Perimeter: perimeter,
	})

	return rectangle, nil
}

// GetRectangle gets a rectangle by ID.
func (rb *RectangleBoundary) GetRectangle(input GetRectangleInput) (*entity.Rectangle, error) {
	rectangle, err := rb.input.GetRectangle(input)
	if err != nil {
		rb.output.ErrorRectangle(err)
		return nil, err
	}

	area := CalculateRectangleArea(rectangle)
	perimeter := CalculateRectanglePerimeter(rectangle)

	rb.output.SuccessRectangle(SuccessRectangleOutput{
		Rectangle: rectangle,
		Area:      area,
		Perimeter: perimeter,
	})

	return rectangle, nil
}

// UpdateRectangle updates a rectangle.
func (rb *RectangleBoundary) UpdateRectangle(input UpdateRectangleInput) (*entity.Rectangle, error) {
	rectangle, err := rb.input.UpdateRectangle(input)
	if err != nil {
		rb.output.ErrorRectangle(err)
		return nil, err
	}

	area := CalculateRectangleArea(rectangle)
	perimeter := CalculateRectanglePerimeter(rectangle)

	rb.output.SuccessRectangle(SuccessRectangleOutput{
		Rectangle: rectangle,
		Area:      area,
		Perimeter: perimeter,
	})

	return rectangle, nil
}

// DeleteRectangle deletes a rectangle by ID.
func (rb *RectangleBoundary) DeleteRectangle(input DeleteRectangleInput) error {
	err := rb.input.DeleteRectangle(input)
	if err != nil {
		rb.output.ErrorRectangle(err)
		return err
	}

	rb.output.SuccessRectangle(SuccessRectangleOutput{})

	return nil
}

// CalculateRectangleArea calculates the area of the rectangle.
func CalculateRectangleArea(rectangle *entity.Rectangle) float64 {
	return float64(rectangle.Width() * rectangle.Height())
}

// CalculateRectanglePerimeter calculates the perimeter of the rectangle.
func CalculateRectanglePerimeter(rectangle *entity.Rectangle) float64 {
	return float64(2 * (rectangle.Width() + rectangle.Height()))
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
