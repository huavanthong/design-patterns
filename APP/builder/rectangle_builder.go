package builder

import (
	"github.com/huavanthong/design-patterns/APP/common"
	"github.com/huavanthong/design-patterns/APP/entity"
)

type IRectangleBuilder interface {
	SetName(name string) IRectangleBuilder
	SetDimensions(dimensions common.Dimensions) IRectangleBuilder
	SetColor(color common.Color) IRectangleBuilder
	SetPosition(position common.Position) IRectangleBuilder
	SetBorderSize(borderSize int) IRectangleBuilder
	Build() entity.Rectangle
}

type RectangleBuilder struct {
	name       string
	dimensions common.Dimensions
	color      common.Color
	position   common.Position
	borderSize int
}

// Constructor
// Ta sẽ đặt tên cho Rectangle ngay lập tức khi ta tạo ra nó.
func NewRectangleBuilder(name string) *RectangleBuilder {
	return &RectangleBuilder{
		name:       name,
		dimensions: common.Dimensions{},
		color:      common.Color{},
		position:   common.Position{},
		borderSize: 1,
	}
}

// Về cơ bản đây không phải là các setter method mà mình biết trong Java.

// Các phương thức SetDimensions(), SetColor() và SetPosition()
// sẽ trả về chính builder struct để có thể gọi tiếp các phương thức khác
// hoặc để kết hợp các phương thức lại với nhau.
func (rb *RectangleBuilder) SetName(name string) IRectangleBuilder {
	rb.name = name
	return rb
}

func (rb *RectangleBuilder) SetDimensions(dimensions common.Dimensions) IRectangleBuilder {
	rb.dimensions.Width = dimensions.Width
	rb.dimensions.Height = dimensions.Height
	return rb
}

func (rb *RectangleBuilder) SetPosition(position common.Position) IRectangleBuilder {
	rb.position = position
	return rb
}

func (rb *RectangleBuilder) SetColor(color common.Color) IRectangleBuilder {
	rb.color = color
	return rb
}

func (rb *RectangleBuilder) SetBorderSize(borderSize int) IRectangleBuilder {
	rb.borderSize = borderSize
	return rb
}

// Build: Combine all setting to build a object
func (rb *RectangleBuilder) Build() entity.Rectangle {
	rectangle := entity.Rectangle{
		Dimensions: rb.dimensions,
		Color:      rb.color,
		Position:   rb.position,
		BorderSize: rb.borderSize,
	}
	return rectangle
}
