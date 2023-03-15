package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/huavanthong/design-patterns/APP/entity"
)

/************************************ Repository in memory ******************************************/
// InMemoryRectangleRepository is an implementation of RectangleRepository that uses an in-memory store.
type InMemoryRectangleRepository struct {
	store map[int]*entity.Rectangle
}

// NewInMemoryRectangleRepository creates a new instance of InMemoryRectangleRepository.
func NewInMemoryRectangleRepository() *InMemoryRectangleRepository {
	return &InMemoryRectangleRepository{
		store: make(map[int]*entity.Rectangle),
	}
}

/**************************************************************************************************************************
Experience 2: Việc ta xây dựng tất cả method trong một struct được gọi là gì?

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
**************************************************************************************************************************/
// Save stores the given rectangle.
// func (r *InMemoryRectangleRepository) Save(rectangle *entity.Rectangle) error {
// 	if _, ok := r.store[rectangle.ID]; ok {
// 		return errors.New("rectangle already exists")
// 	}
// 	r.store[rectangle.ID] = rectangle
// 	return nil
// }

// // GetByID retrieves a rectangle by ID.
// func (r *InMemoryRectangleRepository) GetByID(id int) (*entity.Rectangle, error) {
// 	if rectangle, ok := r.store[id]; ok {
// 		return rectangle, nil
// 	}
// 	return nil, errors.New("rectangle not found")
// }

// // Update updates the given rectangle.
// func (r *InMemoryRectangleRepository) Update(rectangle *entity.Rectangle) error {
// 	if _, ok := r.store[rectangle.ID]; !ok {
// 		return errors.New("rectangle not found")
// 	}
// 	r.store[rectangle.ID] = rectangle
// 	return nil
// }

// // Delete deletes a rectangle by ID.
// func (r *InMemoryRectangleRepository) Delete(id int) error {
// 	if _, ok := r.store[id]; !ok {
// 		return errors.New("rectangle not found")
// 	}
// 	delete(r.store, id)
// 	return nil
// }

/************************************ Repository PostgreSQL ******************************************/
type RectangleRepository struct {
	mu         sync.Mutex
	rectangles []entity.Rectangle
}

func NewRectangleRepository() *RectangleRepository {
	return &RectangleRepository{
		rectangles: []entity.Rectangle{},
	}
}

/*
Experience 1:

Khi sử dụng func (r *RectangleRepository) Save(rectangle entity.Rectangle),
Go sẽ sao chép toàn bộ giá trị của entity.Rectangle vào trong một biến mới trong hàm,
điều này sẽ tốn thêm bộ nhớ để sao chép và giảm hiệu năng của chương trình.

Do đó, chúng ta nên sử dụng
	==> func (r *RectangleRepository) Save(rectangle *entity.Rectangle)
thay vì
	==> func (r *RectangleRepository) Save(rectangle entity.Rectangle)

để tránh tốn thêm bộ nhớ cho việc sao chép. Nếu sử dụng con trỏ, chúng ta chỉ cần truyền địa chỉ của biến đối tượng entity.
Rectangle vào hàm, và các thay đổi trong hàm sẽ được thể hiện trực tiếp trên đối tượng ban đầu mà không cần sao chép toàn bộ giá trị của đối tượng đó.
*/
func (r *RectangleRepository) Save(rectangle *entity.Rectangle) error {

	// Experience 3:
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

func (r *RectangleRepository) FindByID(ID string) (*entity.Rectangle, error) {
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

func (r *RectangleRepository) FindAll() ([]entity.Rectangle, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Return a copy of the slice to avoid race conditions
	rectangles := make([]entity.Rectangle, len(r.rectangles))
	copy(rectangles, r.rectangles)

	return rectangles, nil
}

// Update updates the given rectangle.
func (r *RectangleRepository) Update(rectangle *entity.Rectangle) error {
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

func (r *RectangleRepository) DeleteByID(ID string) error {
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

func (r *RectangleRepository) DeleteAll() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Clear the slice
	r.rectangles = []entity.Rectangle{}
	return nil
}

func (r *RectangleRepository) FindByDimensions(width, height float64) ([]entity.Rectangle, error) {
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

func (r *RectangleRepository) FindByCreatedAt(start, end time.Time) ([]entity.Rectangle, error) {
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
