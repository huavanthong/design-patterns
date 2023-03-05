package boundary

import (
	"github.com/huavanthong/design-patterns/APP/entity"
)

// RectangleInput defines the input port for Rectangle.
type RectangleInput interface {
	CreateRectangle(input CreateRectangleInput) (*entity.Rectangle, error)
	GetRectangle(input GetRectangleInput) (*entity.Rectangle, error)
	UpdateRectangle(input UpdateRectangleInput) (*entity.Rectangle, error)
	DeleteRectangle(input DeleteRectangleInput) error
}

// CreateRectangleInput defines the input structure for creating a rectangle.
type CreateRectangleInput struct {
	Name     string
	Width    int
	Height   int
	Position Position
	Color    Color
}

// GetRectangleInput defines the input structure for getting a rectangle.
type GetRectangleInput struct {
	ID string
}

// UpdateRectangleInput defines the input structure for updating a rectangle.
type UpdateRectangleInput struct {
	ID       string
	Name     string
	Width    int
	Height   int
	Position Position
	Color    Color
}

// DeleteRectangleInput defines the input structure for deleting a rectangle.
type DeleteRectangleInput struct {
	ID string
}
