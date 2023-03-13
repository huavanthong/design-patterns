package validator

import (
	"errors"

	bio "github.com/huavanthong/design-patterns/APP/boundary/io"
	"github.com/huavanthong/design-patterns/APP/common"
)

// Validator defines the interface for validating shapes.
type Validator interface {
	ValidateCreateRectangle(r bio.CreateRectangleInput) error
	ValidateUpdateRectangle(r bio.UpdateRectangleInput) error
	ValidateCreateCircle(c bio.CircleInput) error
}

// ShapeValidator implements the Validator interface for both Rectangle and Circle.
type ShapeValidator struct{}

// NewShapeValidator creates a new instance of ShapeValidator.
func NewShapeValidator() Validator {
	return &ShapeValidator{}
}

// ValidateCreateRectangle validates a Rectangle.
func (v *ShapeValidator) ValidateCreateRectangle(r bio.CreateRectangleInput) error {

	if r.ObjectName == "" {
		return errors.New("object name is required")
	}
	if r.OwnerName == "" {
		return errors.New("owner name is required")
	}
	if r.Dimensions.Width <= 0 {
		return common.ErrInvalidWidth
	}
	if r.Dimensions.Height <= 0 {
		return common.ErrInvalidHeight
	}
	return nil
}

// ValidateUpdateRectangle validates a Rectangle.
func (v *ShapeValidator) ValidateUpdateRectangle(r bio.UpdateRectangleInput) error {
	if r.Dimensions.Width <= 0 {
		return common.ErrInvalidWidth
	}
	if r.Dimensions.Height <= 0 {
		return common.ErrInvalidHeight
	}
	return nil
}

// ValidateCircle validates a Circle.
func (v *ShapeValidator) ValidateCreateCircle(c bio.CircleInput) error {
	if c.Radius <= 0 {
		return common.ErrInvalidRadius
	}
	return nil
}
