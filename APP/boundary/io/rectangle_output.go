package boundary_input_output

import (
	"github.com/huavanthong/design-patterns/APP/common"
	"github.com/huavanthong/design-patterns/APP/entity"
)

// RectangleOutputBoundary defines the output port for Rectangle.
type IRectangleOutput interface {
	SuccessRectangle(output RectangleOutput)
	ErrorRectangle(err error)
}

// RectangleOutput defines the output struct for Rectangle.
type RectangleOutput struct {
	ID         string          `json:"id"`
	ObjectName string          `json:"object_name"`
	Width      float64         `json:"length"`
	Height     float64         `json:"breadth"`
	Position   common.Position `json:"position"`
	Color      common.Color    `json:"color"`
	CreatedAt  string          `json:"created_at"`
	UpdatedAt  string          `json:"updated_at"`
}

/*
Experience 4:
Chúng ta nên xây dựng hàm ToOutput với kiểu trả về là con trỏ tới RectangleOutput, tức là:
	| func ToOutput(r *entity.Rectangle) *RectangleOutput
	thay vì
	| func ToOutput(r *entity.Rectangle) RectangleOutput

Lý do là khi chúng ta truyền vào một con trỏ *entity.Rectangle, ta muốn thay đổi trực tiếp giá trị của biến *RectangleOutput mà không tạo ra một bản sao mới.
Nếu trả về một giá trị thực tế là bản sao của RectangleOutput,
chúng ta sẽ không thể thay đổi giá trị của RectangleOutput ban đầu.
Tuy nhiên, nếu chúng ta trả về con trỏ *RectangleOutput, ta có thể thay đổi giá trị của biến *RectangleOutput mà không tạo ra một bản sao mới.
*/
// ToOutput converts a Rectangle entity to a RectangleOutput struct.
func ToOutput(r *entity.Rectangle) RectangleOutput {
	return RectangleOutput{
		ID:         r.ID,
		ObjectName: r.ObjectName,
		Width:      r.Width,
		Height:     r.Height,
		Position:   r.Position,
		Color:      r.Color,
		CreatedAt:  r.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  r.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// SuccessRectangleOutput defines the success output structure for Rectangle.
type SuccessRectangleOutput struct {
	Rectangle *entity.Rectangle
	Area      float64
	Perimeter float64
}

// ErrorRectangleOutput defines the error output structure for Rectangle.
type ErrorRectangleOutput struct {
	Error string
}
