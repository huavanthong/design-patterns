package repository

import (
	"github.com/huavanthong/design-patterns/APP/entity"
)

// CircleRepository defines the interface for storing and retrieving circles.
/*
Experience design principle 1: Apply Database Agnostic

Nguyên tắc này chỉ ra rằng, mỗi entity chỉ nên define ra các CRUD cơ bản nhất cho chính nó.
Bởi vì việc xây dựng các CRUD cơ bản này giúp chúng ta thực hiện các thao tác cơ bản trên
cơ sở dữ liệu của các hình học một cách đơn giản và linh hoạt. Bất kỳ thay đổi nào về cấu
trúc cơ sở dữ liệu cũng không ảnh hưởng đến khả năng sử dụng của các CRUD này.

==> Từ đó, chúng ta có thể dễ dàng apply nguyên tắc SOLID: Open Close Principle
*/

type CircleRepository interface {
	Create(circle *entity.Circle) error
	GetByID(id string) (*entity.Circle, error)
	Update(circle *entity.Circle) error
	DeleteByID(id string) error
}
