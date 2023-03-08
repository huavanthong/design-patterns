package builder

import (
	"github.com/huavanthong/design-patterns/APP/common"
	"github.com/huavanthong/design-patterns/APP/entity"
)

//
type ICircleBuilder interface {
	SetName(name string) ICircleBuilder
	SetDimensions(dimensions common.Dimensions) ICircleBuilder
	SetColor(color common.Color) ICircleBuilder
	SetPosition(position common.Position) ICircleBuilder
	SetBorderSize(borderSize int) ICircleBuilder
	Build() entity.Circle
}

type CircleBuilder struct {
	name       string
	dimensions common.Dimensions
	color      common.Color
	position   common.Position
	borderSize int
}

// Constructor
// Ta sẽ đặt tên cho Cirle ngay lập tức khi ta tạo ra nó.
func NewCircleBuilder(name string) *CircleBuilder {
	return &CircleBuilder{
		name:       name,
		dimensions: common.Dimensions{},
		color:      common.Color{},
		position:   common.Position{},
		borderSize: 1,
	}
}

// Về cở bản đây không phải là các setter method mà mình biết trong Java.

// Các phương thức SetDimensions(), SetColor() và SetPosition()
// sẽ trả về chính builder struct để có thể gọi tiếp các phương thức khác
// hoặc để kết hợp các phương thức lại với nhau.
func (cb *CircleBuilder) SetName(name string) ICircleBuilder {
	cb.name = name
	return cb
}

func (cb *CircleBuilder) SetDimensions(dimensions common.Dimensions) ICircleBuilder {
	cb.dimensions = dimensions
	return cb
}

func (cb *CircleBuilder) SetColor(color common.Color) ICircleBuilder {
	cb.color = color
	return cb
}

func (cb *CircleBuilder) SetPosition(position common.Position) ICircleBuilder {
	cb.position = position
	return cb
}

func (cb *CircleBuilder) SetBorderSize(borderSize int) ICircleBuilder {
	cb.borderSize = borderSize
	return cb
}

// Build: Combine all setting to build a object
func (cb *CircleBuilder) Build() entity.Circle {
	circle := entity.Circle{
		Dimensions: cb.dimensions,
		Color:      cb.color,
		Position:   cb.position,
		BorderSize: cb.borderSize,
	}
	return circle
}
