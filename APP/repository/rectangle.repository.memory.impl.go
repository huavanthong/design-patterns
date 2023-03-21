package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/huavanthong/design-patterns/APP/entity"
)

/**************************************************************************************************************************
Experience 2: Implement all interface

Trong Go, việc implement một interface được thực hiện thông qua việc cài đặt các phương thức trong một struct (hoặc bất kỳ
kiểu dữ liệu nào khác) và đảm bảo rằng tên và kiểu của các phương thức phải khớp với interface đó.

Trong trường hợp của bạn, interface RectangleRepository được định nghĩa với 4 phương thức:
	+ Save,
 	+ GetByID,
	+ Update,
	+ và Delete.
Bạn đã cài đặt tất cả các phương thức này trong struct InMemoryRectangleRepository. Do đó, InMemoryRectangleRepository đang
thực hiện interface RectangleRepository. Việc thực hiện interface này có nghĩa là bạn đã cung cấp các phương thức cần thiết
để lưu trữ và truy xuất các đối tượng hình chữ nhật trong bộ nhớ.

Khi bạn sử dụng một đối tượng của InMemoryRectangleRepository trong mã của mình, bạn có thể đối xử với nó như với một
RectangleRepository, và gọi các phương thức Save, GetByID, Update, và Delete như bạn mong muốn.

/************************************ Repository PostgreSQL ******************************************/
type InMemoryRectangleRepository struct {
	mu         sync.Mutex
	rectangles []entity.Rectangle
}

func NewInMemoryRectangleRepository() *InMemoryRectangleRepository {
	return &InMemoryRectangleRepository{
		rectangles: []entity.Rectangle{},
	}
}

/*
Experience 1: Experience in using pointer and variable in Go

Khi sử dụng func (r *InMemoryRectangleRepository) Save(rectangle entity.Rectangle),
Go sẽ sao chép toàn bộ giá trị của entity.Rectangle vào trong một biến mới trong hàm,
điều này sẽ tốn thêm bộ nhớ để sao chép và giảm hiệu năng của chương trình.

Do đó, chúng ta nên sử dụng
	==> func (r *InMemoryRectangleRepository) Save(rectangle *entity.Rectangle)
thay vì
	==> func (r *InMemoryRectangleRepository) Save(rectangle entity.Rectangle)

để tránh tốn thêm bộ nhớ cho việc sao chép. Nếu sử dụng con trỏ, chúng ta chỉ cần truyền địa chỉ của biến đối tượng entity.
Rectangle vào hàm, và các thay đổi trong hàm sẽ được thể hiện trực tiếp trên đối tượng ban đầu mà không cần sao chép toàn bộ giá trị của đối tượng đó.
*/
func (r *InMemoryRectangleRepository) Create(rectangle *entity.Rectangle) error {

	// Experience 3: Synchronize data
	// Hàm Save của bạn sử dụng con trỏ đến entity.Rectangle để tránh tạo ra một bản sao không cần thiết của đối tượng,
	// giúp giảm thiểu lượng bộ nhớ sử dụng.
	// Hơn nữa, bạn cũng đã sử dụng lock để đảm bảo tính toàn vẹn dữ liệu trong quá trình thực thi của hàm.
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if the rectangle already exists in the slice
	for i, rect := range r.rectangles {
		if rect.ID == rectangle.ID {
			r.rectangles[i] = *rectangle
			return nil
		}
	}

	// Add the rectangle to the slice
	r.rectangles = append(r.rectangles, *rectangle)
	return nil
}

func (r *InMemoryRectangleRepository) GetByID(ID string) (*entity.Rectangle, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Find the rectangle with the matching ID in the slice
	for _, rect := range r.rectangles {
		if rect.ID == ID {
			return &rect, nil
		}
	}

	return nil, errors.New("rectangle not found")
}

// Update updates the given rectangle.
func (r *InMemoryRectangleRepository) Update(rectangle *entity.Rectangle) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if the rectangle already exists in the slice
	for i, rect := range r.rectangles {
		if rect.ID == rectangle.ID {
			r.rectangles[i] = *rectangle
			return nil
		}
	}

	return errors.New("rectangle not found")
}

func (r *InMemoryRectangleRepository) DeleteByID(ID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Find the index of the rectangle with the matching ID in the slice
	index := -1
	for i, rect := range r.rectangles {
		if rect.ID == ID {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("rectangle not found")
	}

	// Remove the rectangle from the slice
	r.rectangles = append(r.rectangles[:index], r.rectangles[index+1:]...)
	return nil
}

func (r *InMemoryRectangleRepository) FindAll() ([]entity.Rectangle, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Return a copy of the slice to avoid race conditions
	rectangles := make([]entity.Rectangle, len(r.rectangles))
	copy(rectangles, r.rectangles)

	return rectangles, nil
}

func (r *InMemoryRectangleRepository) DeleteAll() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Clear the slice
	r.rectangles = []entity.Rectangle{}
	return nil
}

func (r *InMemoryRectangleRepository) FindByDimensions(width, height float64) ([]entity.Rectangle, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Find all rectangles with matching dimensions in the slice
	var rectangles []entity.Rectangle
	for _, rect := range r.rectangles {
		if rect.Width == width && rect.Height == height {
			rectangles = append(rectangles, rect)
		}
	}

	return rectangles, nil
}

func (r *InMemoryRectangleRepository) FindByCreatedAt(start, end time.Time) ([]entity.Rectangle, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Find all rectangles created between start and end times in the slice
	var rectangles []entity.Rectangle
	for _, rect := range r.rectangles {
		if rect.CreatedAt.After(start) && rect.CreatedAt.Before(end) {
			rectangles = append(rectangles, rect)
		}
	}

	return rectangles, nil
}
