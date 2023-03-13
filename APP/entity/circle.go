// circle.go
package entity

import (
	"fmt"
	"math"
	"time"

	"github.com/huavanthong/design-patterns/APP/common"
)

/*
Các entity được sử dụng để đại diện cho các đối tượng trong hệ thống, thường được sử dụng trong domain logic.
Khi thiết kế entity, nên tập trung vào các đặc tính cốt lõi của đối tượng mà nó đại diện cho.
==> Điều này giúp cho việc hiểu và bảo trì code dễ dàng hơn.

Tuy nhiên, các entity không nên chứa quá nhiều logic.
Nếu entity quá phức tạp hoặc có quá nhiều chức năng, sẽ làm cho code trở nên khó hiểu và khó bảo trì.
Thay vào đó, các entity nên tập trung vào các thuộc tính và phương thức liên quan đến đối tượng mà nó đại diện cho.

Nếu muốn thực hiện các xử lý phức tạp hơn, có thể sử dụng các lớp khác như service hoặc use case để thực hiện.
Vì vậy, việc thiết kế entity nên tập trung vào việc đại diện cho các đối tượng và các thuộc tính của chúng,
và tránh bao gồm quá nhiều logic trong entity.

Update:
	* Đổi Dimensions thành Radius để tập trung vào cốt lõi của entity.const

Refer:
	* https://www.meisternote.com/app/note/vvobbMh1NMEZ/cach-d-t-ten-va-data-type

CommitID:
	*

*/
type Circle struct {
	ID         string          `json:"id"`
	ObjectName string          `json:"object_name"`
	OwnerName  string          `json:"owner_name"`
	Radius     float64         `json:"radius"`
	Color      common.Color    `json:"color"`
	Position   common.Position `json:"position"`
	BorderSize int             `json:"border_size"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}

// utilities
func (c Circle) GetArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) GetPerimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Draw() string {
	return fmt.Sprintf("Drawing Circle with radius %v at position (%v, %v)", c.Radius, c.Position.X, c.Position.Y)
}
