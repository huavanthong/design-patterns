package interactor

import (
	"github.com/huavanthong/design-patterns/APP/boundary"
	"github.com/huavanthong/design-patterns/APP/entity"
	"github.com/huavanthong/design-patterns/APP/validator"
)

// RectangleInteractor là interface cho các nghiệp vụ liên quan đến Rectangle
// RectangleInteractor là interface cho các nghiệp vụ liên quan đến Rectangle
type IRectangleInteractor interface {
	Create(input boundary.CreateRectangleInput) (*entity.Rectangle, error)
	Get(input boundary.GetRectangleInput) (*entity.Rectangle, error)
	Update(input boundary.UpdateRectangleInput) (*entity.Rectangle, error)
	Delete(input boundary.DeleteRectangleInput) error
}

// RectangleInteractor defines the interactor for Rectangle.
type RectangleInteractor struct {
	boundary  boundary.RectangleBoundary
	validator validator.Validator
}

// NewRectangleInteractor creates a new instance of RectangleInteractor.
func NewRectangleInteractor(boundary boundary.RectangleBoundary, validator validator.Validator) *RectangleInteractor {
	return &RectangleInteractor{
		boundary:  boundary,
		validator: validator,
	}
}

// Create creates a rectangle.
func (ri *RectangleInteractor) Create(input boundary.CreateRectangleInput) (*entity.Rectangle, error) {
	// Validate input
	if err := ri.validator.ValidateCreateRectangle(input); err != nil {
		return nil, err
	}

	// Create rectangle
	return ri.boundary.CreateRectangle(input)
}

// Get gets a rectangle by ID.
func (ri *RectangleInteractor) Get(input boundary.GetRectangleInput) (*entity.Rectangle, error) {
	return ri.boundary.GetRectangle(input)
}

// Update updates a rectangle.
func (ri *RectangleInteractor) Update(input boundary.UpdateRectangleInput) (*entity.Rectangle, error) {
	// Validate input
	if err := ri.validator.ValidateUpdateRectangle(input); err != nil {
		return nil, err
	}

	// Update rectangle
	return ri.boundary.UpdateRectangle(input)
}

// Delete deletes a rectangle by ID.
func (ri *RectangleInteractor) Delete(input boundary.DeleteRectangleInput) error {
	return ri.boundary.DeleteRectangle(input)
}
