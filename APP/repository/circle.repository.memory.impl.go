package repository

import (
	"errors"

	"github.com/huavanthong/design-patterns/APP/entity"
)

// InMemoryCircleRepository is an implementation of RectangleRepository that uses an in-memory store.
type InMemoryCircleRepository struct {
	store map[string]*entity.Circle
}

// NewInMemoryCircleRepository creates a new instance of InMemoryCircleRepository.
func NewInMemoryCircleRepository() *InMemoryCircleRepository {
	return &InMemoryCircleRepository{
		store: make(map[string]*entity.Circle),
	}
}

// Save stores the given circle.
func (r *InMemoryCircleRepository) Save(circle *entity.Circle) error {
	if _, ok := r.store[circle.ID]; ok {
		return errors.New("circle already exists")
	}
	r.store[circle.ID] = circle
	return nil
}

// GetByID retrieves a circle by ID.
func (r *InMemoryCircleRepository) GetByID(id string) (*entity.Circle, error) {
	if circle, ok := r.store[id]; ok {
		return circle, nil
	}
	return nil, errors.New("circle not found")
}

// Update updates the given circle.
func (r *InMemoryCircleRepository) Update(circle *entity.Circle) error {
	if _, ok := r.store[circle.ID]; !ok {
		return errors.New("circle not found")
	}
	r.store[circle.ID] = circle
	return nil
}

// Delete deletes a circle by ID.
func (r *InMemoryCircleRepository) Delete(id string) error {
	if _, ok := r.store[id]; !ok {
		return errors.New("circle not found")
	}
	delete(r.store, id)
	return nil
}
