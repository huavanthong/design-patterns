package repository

import (
	"github.com/huavanthong/design-patterns/APP/entity"
)

// CircleRepository defines the interface for storing and retrieving circles.
type CircleRepository interface {
	Save(circle *entity.Circle) error
	GetByID(id int) (*entity.Circle, error)
	Update(circle *entity.Circle) error
	Delete(id int) error
}
