package interactor

import "github.com/huavanthong/design-patterns/APP/entity"

// Validator defines the interface for validating shapes.
type Validator interface {
	ValidateRectangle(r *entity.Rectangle) error
	ValidateCircle(c *entity.Circle) error
}

// ShapeValidator implements the Validator interface for both Rectangle and Circle.
type ShapeValidator struct{}

// ValidateRectangle validates a Rectangle.
func (v *ShapeValidator) ValidateRectangle(r *entity.Rectangle) error {
	if r.Width <= 0 {
		return ErrInvalidWidth
	}
	if r.Height <= 0 {
		return ErrInvalidHeight
	}
	return nil
}

// ValidateCircle validates a Circle.
func (v *ShapeValidator) ValidateCircle(c *entity.Circle) error {
	if c.Radius <= 0 {
		return ErrInvalidRadius
	}
	return nil
}
