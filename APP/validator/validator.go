package interactor

import (
	"github.com/huavanthong/design-patterns/APP/boundary"
	"github.com/huavanthong/design-patterns/APP/common"
)

// Validator defines the interface for validating shapes.
type Validator interface {
	ValidateCreateRectangle(r boundary.CreateRectangleInput) error
	ValidateUpdateRectangle(r boundary.UpdateRectangleInput) error
	ValidateCreateCircle(c boundary.CircleInput) error
}

// ShapeValidator implements the Validator interface for both Rectangle and Circle.
type ShapeValidator struct{}

// NewShapeValidator creates a new instance of ShapeValidator.
func NewShapeValidator() Validator {
	return &ShapeValidator{}
}

// ValidateCreateRectangle validates a Rectangle.
func (v *ShapeValidator) ValidateCreateRectangle(r boundary.CreateRectangleInput) error {
	if r.Width <= 0 {
		return common.ErrInvalidWidth
	}
	if r.Height <= 0 {
		return common.ErrInvalidHeight
	}
	return nil
}

// ValidateUpdateRectangle validates a Rectangle.
func (v *ShapeValidator) ValidateUpdateRectangle(r boundary.UpdateRectangleInput) error {
	if r.Width <= 0 {
		return common.ErrInvalidWidth
	}
	if r.Height <= 0 {
		return common.ErrInvalidHeight
	}
	return nil
}

// ValidateCircle validates a Circle.
func (v *ShapeValidator) ValidateCreateCircle(c boundary.CircleInput) error {
	if c.Radius <= 0 {
		return common.ErrInvalidRadius
	}
	return nil
}
