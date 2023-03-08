package boundary

import (
	"errors"
	"fmt"

	"github.com/huavanthong/design-patterns/APP/entity"
	"github.com/huavanthong/design-patterns/APP/interactor"
	"github.com/huavanthong/design-patterns/APP/validator"
)

// RectangleBoundary defines the boundary context for Rectangle.
type RectangleBoundary struct {
	interactor interactor.IRectangleInteractor
	validator  *validator.Validator
}

// NewRectangleBoundary is a factory function
// that it returns a new instance of RectangleBoundary.
func NewRectangleBoundary(interactor interactor.IRectangleInteractor, validator *validator.Validator) *RectangleBoundary {
	return &RectangleBoundary{
		interactor: interactor,
		validator:  validator,
	}
}

// CreateRectangle creates a rectangle.
func (rb *RectangleBoundary) CreateRectangle(input CreateRectangleInput) (*RectangleOutput, error) {
	if err := rb.validator.ValidateCreateInput(input); err != nil {
		return nil, err
	}

	rectangle := input.ToEntity()
	err := rb.interactor.CreateRectangle(rectangle)
	if err != nil {
		return nil, err
	}

	output := RectangleOutput{
		ID:        rectangle.ID,
		Name:      rectangle.Name,
		Length:    rectangle.Length,
		Breadth:   rectangle.Breadth,
		CreatedAt: rectangle.CreatedAt,
		UpdatedAt: rectangle.UpdatedAt,
	}
	return &output, nil
}

// GetRectangle gets a rectangle by ID.
func (rb *RectangleBoundary) GetRectangle(input GetRectangleInput) (*RectangleOutput, error) {
	if input.ID == "" {
		return nil, errors.New("ID is required")
	}

	rectangle, err := rb.interactor.GetRectangle(input.ID)
	if err != nil {
		return nil, err
	}

	output := RectangleOutput{
		ID:        rectangle.ID,
		Name:      rectangle.Name,
		Length:    rectangle.Length,
		Breadth:   rectangle.Breadth,
		CreatedAt: rectangle.CreatedAt,
		UpdatedAt: rectangle.UpdatedAt,
	}

	return &output, nil
}

// UpdateRectangle updates a rectangle.
func (rb *RectangleBoundary) UpdateRectangle(input UpdateRectangleInput) (*RectangleOutput, error) {

	rectangle, err := rb.interactor.UpdateRectangle(input.ID)
	if err != nil {
		return nil, err
	}

	output := RectangleOutput{
		ID:        rectangle.ID,
		Name:      rectangle.Name,
		Length:    rectangle.Length,
		Breadth:   rectangle.Breadth,
		CreatedAt: rectangle.CreatedAt,
		UpdatedAt: rectangle.UpdatedAt,
	}

	return &output, nil
}

// DeleteRectangle deletes a rectangle by ID.
func (rb *RectangleBoundary) DeleteRectangle(input DeleteRectangleInput) error {
	err := rb.interactor.DeleteRectangle(input)
	if err != nil {
		return err
	}
	return nil
}

// CalculateRectangleArea calculates the area of the rectangle.
func CalculateRectangleArea(rectangle *entity.Rectangle) float64 {
	return float64(rectangle.Dimensions.Width * rectangle.Dimensions.Height)
}

// CalculateRectanglePerimeter calculates the perimeter of the rectangle.
func CalculateRectanglePerimeter(rectangle *entity.Rectangle) float64 {
	return float64(2 * (rectangle.Dimensions.Width + rectangle.Dimensions.Height))
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
