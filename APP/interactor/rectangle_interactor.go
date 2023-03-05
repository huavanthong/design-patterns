package interactor

import (
	"github.com/huavanthong/design-patterns/APP/boundary"
	"github.com/huavanthong/design-patterns/APP/entity"
)

// RectangleInteractor defines the interactor for Rectangle.
type RectangleInteractor struct {
	boundary  boundary.RectangleBoundary
	validator *Validator
}

// NewRectangleInteractor creates a new instance of RectangleInteractor.
func NewRectangleInteractor(boundary boundary.RectangleBoundary, validator *Validator) *RectangleInteractor {
	return &RectangleInteractor{
		boundary:  boundary,
		validator: validator,
	}
}

// CreateRectangle creates a new Rectangle entity with the provided input.
func (i *RectangleInteractor) CreateRectangle(input boundary.RectangleInput) {
	if err := i.validator.Validate(input); err != nil {
		i.boundary.ErrorRectangle(err)
		return
	}

	rectangle := entity.NewRectangle(input.Width, input.Height)

	if err := rectangle.Validate(); err != nil {
		i.boundary.ErrorRectangle(err)
		return
	}

	output := boundary.SuccessRectangleOutput{
		Rectangle: rectangle,
	}

	i.boundary.SuccessRectangle(output)
}

// GetRectangle retrieves a Rectangle entity with the provided input.
func (i *RectangleInteractor) GetRectangle(input boundary.RectangleInput) {
	if err := i.validator.ValidateRectangle(input); err != nil {
		i.boundary.ErrorRectangle(err)
		return
	}

	rectangle := entity.NewRectangle(input.Width, input.Height)

	if err := rectangle.Validate(); err != nil {
		i.boundary.ErrorRectangle(err)
		return
	}

	area := rectangle.Area()
	perimeter := rectangle.Perimeter()

	output := boundary.RectangleInfoOutput{
		Area:      area,
		Perimeter: perimeter,
	}

	i.boundary.RectangleInfo(output)
}
