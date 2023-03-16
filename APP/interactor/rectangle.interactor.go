package interactor

import (
	"github.com/huavanthong/design-patterns/APP/entity"
)

// RectangleInteractor is an interface for managing rectangles
type IRectangleInteractor interface {
	CreateRectangle(rectangle *entity.Rectangle) error
	GetRectangleByID(id string) (*entity.Rectangle, error)
	UpdateRectangle(rectangle *entity.Rectangle) error
	DeleteRectangleByID(id string) error
}
