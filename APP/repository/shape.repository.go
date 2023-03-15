package repository

import (
	"time"

	"github.com/huavanthong/design-patterns/APP/entity"
)

// ShapeRepository defines the interface for storing and retrieving shapes.
type IShapeRepository interface {
	// Experience 1:
	Save(rectangle *entity.Rectangle)
	FindByID(ID string) (*entity.Rectangle, error)
	FindAll() ([]entity.Rectangle, error)
	Update(rectangle *entity.Rectangle) error
	DeleteByID(ID string) error
	DeleteAll() error
	FindByDimensions(width, hieght float64) ([]entity.Rectangle, error)
	FindByCreatedAt(start, end time.Time) ([]entity.Rectangle, error)
}
