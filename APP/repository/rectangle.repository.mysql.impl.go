package repository

import (
	"time"

	"github.com/huavanthong/design-patterns/APP/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type RectangleRepository interface {
	Save(rectangle *entity.Rectangle) error
	FindByID(ID string) (*entity.Rectangle, error)
	FindAll() ([]entity.Rectangle, error)
	Update(rectangle *entity.Rectangle) error
	DeleteByID(ID string) error
	DeleteAll() error
	FindByDimensions(width, height float64) ([]entity.Rectangle, error)
	FindByCreatedAt(start, end time.Time) ([]entity.Rectangle, error)
}

type rectangleRepository struct {
	db *gorm.DB
}

func NewRectangleRepository(db *gorm.DB) RectangleRepository {
	return &rectangleRepository{db}
}

func (repo *rectangleRepository) Save(rectangle *entity.Rectangle) error {
	return repo.db.Save(rectangle).Error
}

func (repo *rectangleRepository) FindByID(ID string) (*entity.Rectangle, error) {
	var rectangle entity.Rectangle
	err := repo.db.Where("id = ?", ID).First(&rectangle).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &rectangle, nil
}

func (repo *rectangleRepository) FindAll() ([]entity.Rectangle, error) {
	var rectangles []entity.Rectangle
	err := repo.db.Find(&rectangles).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return rectangles, nil
}

// các method khác cũng tương tự
