// rectangle.go
package entity

import (
	"fmt"
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
	* Đổi Dimensions thành Width, Height để tập trung vào cốt lõi của entity.const

Refer:
	* https://www.meisternote.com/app/note/vvobbMh1NMEZ/cach-d-t-ten-va-data-type

CommitID:
	* cb1da2e8a84a09d7bb2de065da1094ea0067948d
*/
type Rectangle struct {
	ID         string          `json:"id"`
	ObjectName string          `json:"object_name"`
	OwnerName  string          `json:"owner_name"`
	Width      float64         `json:"width"`
	Height     float64         `json:"height"`
	Color      common.Color    `json:"color"`
	Position   common.Position `json:"position"`
	BorderSize int             `json:"border_size"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}

// utilities
func (r Rectangle) GetArea() float64 {
	return r.Width * r.Height
}

func (r Rectangle) GetPerimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Draw() string {
	return fmt.Sprintf("Drawing Rectangle with width %v, height %v at position (%v, %v)", r.Width, r.Height, r.Position.X, r.Position.Y)
}

// Nhờ vào việc xây dựng các Dimensions ta sẽ dàng phát triển đa hình học với ví dụ sau đây.
// Refer:
