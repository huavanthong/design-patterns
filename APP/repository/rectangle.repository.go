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

Experience 7: Declare interface for using the variant database, and apply Open/Close Principle
Conclusion:
	+ Bởi vì chúng ta đang triển khai repository cho cả In memory và MongoDB nên ta quyết định tạo interface cho repository này.
	+

**************************************************************************************************************************/
// RectangleRepository defines the interface for storing and retrieving rectangles.

/* Design Old */
type RectangleRepositoryOld interface {
	// Experience 1:
	Save(rectangle *entity.Rectangle)
	FindByID(ID string) (*entity.Rectangle, error)
	FindAll() ([]entity.Rectangle, error)
	Update(rectangle *entity.Rectangle) error
	DeleteByID(ID string) error
	// Error 1: Don't DeleteAll() for specific entity.
	DeleteAll() error
	FindByDimensions(width, hieght float64) ([]entity.Rectangle, error)
	FindByCreatedAt(start, end time.Time) ([]entity.Rectangle, error)
}

/*
Experience SOLID 1: Open Close Principle

Việc tạo ra interface RectangleSearchRepository để tách các phương thức tìm kiếm vào một interface riêng
là một ví dụ tốt về việc áp dụng nguyên tắc Open/Close Principle.
Khi chúng ta cần thêm các phương thức tìm kiếm mới hoặc thay đổi cách tìm kiếm, chúng ta chỉ cần thay đổi
trong interface này mà không làm ảnh hưởng đến các phương thức khác trong RectangleRepository.

*/
type RectangleSearchRepository interface {
	FindByID(ID string) (*entity.Rectangle, error)
	FindByDimensions(width, height float64) ([]entity.Rectangle, error)
	FindByCreatedAt(start, end time.Time) ([]entity.Rectangle, error)
}

/*
 Experience design principle 1: Apply database agnostic


*/
type RectangleRepository interface {
	// Experience 1: Experience in using pointer and variable
	Create(rectangle *entity.Rectangle) error
	GetByID(id string) ([]entity.Rectangle, error)
	Update(rectangle *entity.Rectangle) error
	DeleteByID(ID string) error
}

/*
Experience SOLID 2: Interface Segregation Principle 4

Việc sử dụng RectangleRepositoryComposite để kết hợp các interface RectangleRepository và RectangleSearchRepository
cũng là một ví dụ tốt về việc áp dụng nguyên tắc Interface Segregation Principle (ISP),
bởi vì chúng ta chỉ kết hợp những phương thức cần thiết và không cần thiết phải kết hợp tất cả các phương thức từ
các interface khác nhau.
*/
type RectangleRepositoryComposite interface {
	RectangleRepository
	RectangleSearchRepository
}
