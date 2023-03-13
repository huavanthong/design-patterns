package boundary

import (
	"errors"
	"fmt"

	bio "github.com/huavanthong/design-patterns/APP/boundary/io"
	"github.com/huavanthong/design-patterns/APP/interactor"

	"github.com/huavanthong/design-patterns/APP/validator"
)

// RectangleBoundary defines the boundary context for Rectangle.
type RectangleBoundary struct {
	interactor interactor.IRectangleInteractor
	validator  validator.ShapeValidator
}

// NewRectangleBoundary is a factory function
// that it returns a new instance of RectangleBoundary.
func NewRectangleBoundary(interactor interactor.IRectangleInteractor, validator validator.ShapeValidator) *RectangleBoundary {
	return &RectangleBoundary{
		interactor: interactor,
		validator:  validator,
	}
}

// CreateRectangle creates a rectangle.
func (rb *RectangleBoundary) CreateRectangle(input bio.CreateRectangleInput) (*bio.RectangleOutput, error) {
	if err := rb.validator.ValidateCreateRectangle(input); err != nil {
		return nil, err
	}

	rectangle := input.ToEntity()
	err := rb.interactor.CreateRectangle(rectangle)
	if err != nil {
		return nil, err
	}

	output := bio.ToOutput(rectangle)

	return &output, nil
}

// GetRectangle gets a rectangle by ID.
func (rb *RectangleBoundary) GetRectangle(input bio.GetRectangleInput) (*bio.RectangleOutput, error) {
	if input.ID == "" {
		return nil, errors.New("ID is required")
	}

	rectangle, err := rb.interactor.GetRectangleByID(input.ID)
	if err != nil {
		return nil, err
	}

	output := bio.ToOutput(*rectangle)

	return &output, nil
}

// UpdateRectangle updates a rectangle.
func (rb *RectangleBoundary) UpdateRectangle(input bio.UpdateRectangleInput) (*bio.RectangleOutput, error) {

	if err := rb.validator.ValidateUpdateRectangle(input); err != nil {
		return nil, err
	}

	rectangle := input.UpdateEntity()
	err := rb.interactor.UpdateRectangle(rectangle)
	if err != nil {
		return nil, err
	}

	output := bio.ToOutput(rectangle)

	return &output, nil
}

// DeleteRectangle deletes a rectangle by ID.
func (rb *RectangleBoundary) DeleteRectangle(input bio.DeleteRectangleInput) error {
	err := rb.interactor.DeleteRectangleByID(input.ID)
	if err != nil {
		return err
	}
	return nil
}

// SuccessRectangle is the success output method.
func (rb *RectangleBoundary) SuccessRectangle(output bio.RectangleOutput) {
	// fmt.Printf("Rectangle: %+v\n", output.Rectangle)
	// fmt.Printf("Area: %f\n", output.Area)
	// fmt.Printf("Perimeter: %f\n", output.Perimeter)
}

// ErrorRectangle is the error output method.
func (rb *RectangleBoundary) ErrorRectangle(err error) {
	fmt.Printf("Error: %s\n", err)
}
