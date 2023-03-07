package builder

import (
	"github.com/huavanthong/design-patterns/APP/common"
	"github.com/huavanthong/design-patterns/APP/entity"
)

type RectangleBuilder interface {
	SetDimensions(dimensions common.Dimensions) RectangleBuilder
	SetColor(color common.Color) RectangleBuilder
	SetPosition(position common.Position) RectangleBuilder
	SetBorderSize(borderSize int) RectangleBuilder
	Build() entity.Rectangle
}

type rectangleBuilder struct {
	dimensions common.Dimensions
	color      common.Color
	position   common.Position
	borderSize int
}

// Về cở bản đây không phải là các setter method mà mình biết trong Java.

// Các phương thức SetDimensions(), SetColor() và SetPosition()
// sẽ trả về chính builder struct để có thể gọi tiếp các phương thức khác
// hoặc để kết hợp các phương thức lại với nhau.
func (rb *rectangleBuilder) SetDimensions(dimensions common.Dimensions) RectangleBuilder {
	rb.dimensions.Width = dimensions.Width
	rb.dimensions.Height = dimensions.Height
	return rb
}

func (rb *rectangleBuilder) SetPosition(position common.Position) RectangleBuilder {
	rb.position = position
	return rb
}

func (rb *rectangleBuilder) SetColor(color common.Color) RectangleBuilder {
	rb.color = color
	return rb
}

func (rb *rectangleBuilder) SetBorderSize(borderSize int) RectangleBuilder {
	rb.borderSize = borderSize
	return rb
}

// Build: Combine all setting to build a object
func (rb *rectangleBuilder) Build() entity.Rectangle {
	rectangle := entity.Rectangle{
		Dimensions: rb.dimensions,
		Color:      rb.color,
		Position:   rb.position,
		BorderSize: rb.borderSize,
	}
	return rectangle
}
