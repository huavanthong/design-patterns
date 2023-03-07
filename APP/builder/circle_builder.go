package builder

import (
	"github.com/huavanthong/design-patterns/APP/common"
	"github.com/huavanthong/design-patterns/APP/entity"
)

//
type CircleBuilder interface {
	SetDimensions(dimensions common.Dimensions) CircleBuilder
	SetColor(color common.Color) CircleBuilder
	SetPosition(position common.Position) CircleBuilder
	SetBorderSize(position common.Position) CircleBuilder
	Build() entity.Circle
}

type circleBuilder struct {
	dimensions common.Dimensions
	color      common.Color
	position   common.Position
	borderSize int
}

// Về cở bản đây không phải là các setter method mà mình biết trong Java.

// Các phương thức SetDimensions(), SetColor() và SetPosition()
// sẽ trả về chính builder struct để có thể gọi tiếp các phương thức khác
// hoặc để kết hợp các phương thức lại với nhau.
func (cb *circleBuilder) SetDimensions(dimensions common.Dimensions) CircleBuilder {
	cb.dimensions = dimensions
	return cb
}

func (cb *circleBuilder) SetColor(color common.Color) CircleBuilder {
	cb.color = color
	return cb
}

func (cb *circleBuilder) SetPosition(position common.Position) CircleBuilder {
	cb.position = position
	return cb
}

func (cb *circleBuilder) SetBorderSize(position common.Position) CircleBuilder {
	cb.position = position
	return cb
}

// Build: Combine all setting to build a object
func (cb *circleBuilder) Build() entity.Circle {
	circle := entity.Circle{
		Dimensions: cb.dimensions,
		Color:      cb.color,
		Position:   cb.position,
		BorderSize: cb.borderSize,
	}
	return circle
}
