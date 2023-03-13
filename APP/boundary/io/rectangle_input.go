package boundary_input_output

import (
	"time"

	"github.com/huavanthong/design-patterns/APP/common"
	"github.com/huavanthong/design-patterns/APP/entity"
)

// RectangleInput defines the input port for Rectangle.
type IRectangleInput interface {
	CreateRectangle(input CreateRectangleInput) (*RectangleOutput, error)
	GetRectangle(input GetRectangleInput) (*RectangleOutput, error)
	UpdateRectangle(input UpdateRectangleInput) (*RectangleOutput, error)
	DeleteRectangle(input DeleteRectangleInput) error
}

// CreateRectangleInput defines the input structure for creating a rectangle.
type CreateRectangleInput struct {
	ObjectName string
	OwnerName  string
	Dimensions common.Dimensions
	Position   common.Position
	Color      common.Color
}

func (ri *CreateRectangleInput) ToEntity() *entity.Rectangle {
	return &entity.Rectangle{
		ID:         common.GenerateUUID(),
		ObjectName: ri.ObjectName,
		OwnerName:  ri.OwnerName,
		Width:      ri.Dimensions.Width,
		Height:     ri.Dimensions.Height,
		Position:   ri.Position,
		Color:      ri.Color,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

// GetRectangleInput defines the input structure for getting a rectangle.
type GetRectangleInput struct {
	ID string
}

// UpdateRectangleInput defines the input structure for updating a rectangle.
type UpdateRectangleInput struct {
	ID         string
	ObjectName string
	OwnerName  string
	Dimensions common.Dimensions
	Position   common.Position
	Color      common.Color
}

func (ri *UpdateRectangleInput) UpdateEntity() *entity.Rectangle {
	return &entity.Rectangle{
		ID:         ri.ID,
		ObjectName: ri.ObjectName,
		OwnerName:  ri.OwnerName,
		Width:      ri.Dimensions.Width,
		Height:     ri.Dimensions.Height,
		Position:   ri.Position,
		Color:      ri.Color,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

// DeleteRectangleInput defines the input structure for deleting a rectangle.
type DeleteRectangleInput struct {
	ID string
}
