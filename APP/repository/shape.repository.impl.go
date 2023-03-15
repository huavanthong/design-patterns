package repository

import (
	"errors"
	"sync"

	"github.com/huavanthong/design-patterns/APP/entity"
)

type ShapeRepository struct {
	mu     sync.RWMutex
	shapes []entity.Shape
	nextID int
}

func NewShapeRepository() *ShapeRepository {
	return &ShapeRepository{
		shapes: make([]entity.Shape, 0),
		nextID: 1,
	}
}

func (sr *ShapeRepository) Save(s entity.Shape) (int, error) {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	if s.GetID() == 0 {
		s.SetID(sr.nextID)
		sr.nextID++
	}

	for i, shape := range sr.shapes {
		if shape.GetID() == s.GetID() {
			sr.shapes[i] = s
			return s.GetID(), nil
		}
	}

	sr.shapes = append(sr.shapes, s)
	return s.GetID(), nil
}

func (sr *ShapeRepository) FindByID(id int) (entity.Shape, error) {
	sr.mu.RLock()
	defer sr.mu.RUnlock()

	for _, shape := range sr.shapes {
		if shape.GetID() == id {
			return shape, nil
		}
	}

	return nil, errors.New("shape not found")
}

func (sr *ShapeRepository) FindAll() []entity.Shape {
	sr.mu.RLock()
	defer sr.mu.RUnlock()

	shapes := make([]entity.Shape, len(sr.shapes))
	copy(shapes, sr.shapes)

	return shapes
}

func (sr *ShapeRepository) DeleteByID(id int) error {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	for i, shape := range sr.shapes {
		if shape.GetID() == id {
			sr.shapes = append(sr.shapes[:i], sr.shapes[i+1:]...)
			return nil
		}
	}

	return errors.New("shape not found")
}
