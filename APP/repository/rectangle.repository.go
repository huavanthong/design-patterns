package repository

import (
	"time"

	"github.com/huavanthong/design-patterns/APP/entity"
)

/**************************************************************************************************************************
Question 1: Ta hỏi ChatGPT có cần tạo interace cho Repository này không?

Nếu muốn đảm bảo tính trừu tượng của code và dễ dàng thay thế implementation thì nên tạo interface cho RectangleRepository.
Tuy nhiên, trong trường hợp này, nếu chỉ sử dụng một implementation cho RectangleRepository (InMemoryRectangleRepository)
và không dự định thay đổi implementation này trong tương lai, thì có thể bỏ qua việc tạo interface và sử dụng trực tiếp
implementation này.
**************************************************************************************************************************/
// RectangleRepository defines the interface for storing and retrieving rectangles.
type IRectangleRepository interface {
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
