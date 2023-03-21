package repository

import (
	"errors"
	"sync"

	"github.com/huavanthong/design-patterns/APP/entity"
)

type InMemoryShapeRepository struct {
	mu     sync.RWMutex
	shapes []entity.Shape
}

func NewShapeRepository() *InMemoryShapeRepository {
	return &InMemoryShapeRepository{
		shapes: []entity.Shape{},
	}
}

// Create stores the given shape.
func (s *InMemoryShapeRepository) Create(shape entity.Shape) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the shape already exists in the slice
	for i, v := range s.shapes {
		if v.ID == shape.ID {
			s.shapes[i] = *shape
			return nil
		}
	}

	// Add the shape to the slice
	s.shapes = append(s.shapes, shape)
	return nil
}

// GetByID retrieves a shape by ID.
func (s *InMemoryShapeRepository) GetByID(ID string) (entity.Shape, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Find the shape with the matching ID in the slice
	for _, v := range s.shapes {
		if v.ID == ID {
			return &v, nil
		}
	}

	return nil, errors.New("shape not found")
}

// Update updates the given shape.
func (s *InMemoryShapeRepository) Update(shape *entity.Shape) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the shape already exists in the slice
	for i, v := range s.shapes {
		if v.ID == shape.ID {
			s.shapes[i] = *shape
			return nil
		}
	}

	return errors.New("shape not found")
}

// DeleteByID deletes a circle by ID.
func (s *InMemoryShapeRepository) DeleteByID(ID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, v := range s.shapes {
		if v.ID == ID {
			s.shapes = append(s.shapes[:i], s.shapes[i+1:]...)
			return nil
		}
	}

	return errors.New("shape not found")
}
