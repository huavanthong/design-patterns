package boundary

import (
	"github.com/huavanthong/design-patterns/APP/entity"
)

// RectangleBoundary is the boundary for Rectangle.
type RectangleBoundary struct {
	interactor interactor.RectangleInteractor
}

// NewRectangleBoundary creates a new RectangleBoundary.
func NewRectangleBoundary(interactor interactor.RectangleInteractor) *RectangleBoundary {
	return &RectangleBoundary{
		interactor: interactor,
	}
}

// CreateRectangle creates a rectangle.
func (rb *RectangleBoundary) CreateRectangle(input CreateRectangleInput) (*entity.Rectangle, error) {
	rectangleBuilder := NewRectangleBuilder()
	rectangleBuilder.SetName(input.Name)
	rectangleBuilder.SetWidth(input.Width)
	rectangleBuilder.SetHeight(input.Height)
	rectangleBuilder.SetPosition(input.Position)
	rectangleBuilder.SetColor(input.Color)

	rectangle, err := rb.interactor.CreateRectangle(rectangleBuilder)
	if err != nil {
		return nil, err
	}

	return rectangle, nil
}

// GetRectangle gets a rectangle.
func (rb *RectangleBoundary) GetRectangle(input GetRectangleInput) (*entity.Rectangle, error) {
	rectangle, err := rb.interactor.GetRectangle(input.ID)
	if err != nil {
		return nil, err
	}

	return rectangle, nil
}

// UpdateRectangle updates a rectangle.
func (rb *RectangleBoundary) UpdateRectangle(input UpdateRectangleInput) (*entity.Rectangle, error) {
	rectangleBuilder := NewRectangleBuilder()
	rectangleBuilder.SetID(input.ID)
	rectangleBuilder.SetName(input.Name)
	rectangleBuilder.SetWidth(input.Width)
	rectangleBuilder.SetHeight(input.Height)
	rectangleBuilder.SetPosition(input.Position)
	rectangleBuilder.SetColor(input.Color)

	rectangle, err := rb.interactor.UpdateRectangle(rectangleBuilder)
	if err != nil {
		return nil, err
	}

	return rectangle, nil
}

// DeleteRectangle deletes a rectangle.
func (rb *RectangleBoundary) DeleteRectangle(input DeleteRectangleInput) error {
	err := rb.interactor.DeleteRectangle(input.ID)
	if err != nil {
		return err
	}

	return nil
}
