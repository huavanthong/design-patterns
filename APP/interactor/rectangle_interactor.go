package interactor

import (
	"github.com/huavanthong/design-patterns/APP/entity"
	"github.com/huavanthong/design-patterns/APP/repository"
)

// RectangleInteractor is an interface for managing rectangles
type IRectangleInteractor interface {
	CreateRectangle(rectangle entity.Rectangle) error
	GetRectangleByID(id string) (*entity.Rectangle, error)
	UpdateRectangle(rectangle entity.Rectangle) error
	DeleteRectangleByID(id string) error
}

type rectangleInteractor struct {
	rectangleRepository repository.RectangleRepository
}

// NewRectangleInteractor returns a new instance of RectangleInteractor
func NewRectangleInteractor(rectangleRepository repository.RectangleRepository) IRectangleInteractor {
	return &rectangleInteractor{
		rectangleRepository: rectangleRepository,
	}
}

// CreateRectangle creates a new rectangle
func (ri *rectangleInteractor) CreateRectangle(rectangle entity.Rectangle) error {
	return ri.rectangleRepository.Save(rectangle)
}

// GetRectangleByID returns a rectangle by ID
func (ri *rectangleInteractor) GetRectangleByID(id string) (*entity.Rectangle, error) {
	return ri.rectangleRepository.FindByID(id)
}

// UpdateRectangle updates an existing rectangle
func (ri *rectangleInteractor) UpdateRectangle(rectangle entity.Rectangle) error {
	return ri.rectangleRepository.Update(rectangle)
}

// DeleteRectangleByID deletes a rectangle by ID
func (ri *rectangleInteractor) DeleteRectangleByID(id string) error {
	return ri.rectangleRepository.DeleteByID(id)
}
