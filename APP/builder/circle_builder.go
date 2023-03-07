package builder

import (
	"github.com/huavanthong/design-patterns/APP/common"
	"github.com/huavanthong/design-patterns/APP/entity"
)

//
type ICircleBuilder interface {
	SetDimensions(dimensions common.Dimensions) ICircleBuilder
	SetColor(color common.Color) ICircleBuilder
	SetPosition(position common.Position) ICircleBuilder
	SetBorderSize(borderSize int) ICircleBuilder
	Build() entity.Circle
}

type CircleBuilder struct {
	dimensions common.Dimensions
	color      common.Color
	position   common.Position
	borderSize int
}

// Constructor
func NewCircleBuilder() *CircleBuilder {
	return &CircleBuilder{
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
