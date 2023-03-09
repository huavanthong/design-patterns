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
	Name     string
	Width    float64
	Height   float64
	Position common.Position
	Color    common.Color
}

func (ri *CreateRectangleInput) ToEntity() *entity.Rectangle {
	return &entity.Rectangle{
		ID:   common.GenerateUUID(),
		Name: ri.Name,
		Dimensions: common.Dimensions{
			Width:  ri.Width,
			Height: ri.Height,
			Radius: 0},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
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
	Position common.Position
	Color    common.Color
}

// DeleteRectangleInput defines the input structure for deleting a rectangle.
type DeleteRectangleInput struct {
	ID string
}
