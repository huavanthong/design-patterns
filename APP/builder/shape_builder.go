// File shape_builder.go
package builder

import (
	"github.com/huavanthong/design-patterns/APP/common"
	"github.com/huavanthong/design-patterns/APP/entity"
)

type IShapeBuilder interface {
	SetDimensions(dimensions common.Dimensions) IShapeBuilder
	SetColor(color common.Color) IShapeBuilder
	SetPosition(position common.Position) IShapeBuilder
	SetBorderSize(borderSize int) IShapeBuilder
	Build() entity.Shape
}

// Mẫu thiết kế mới sử dụng kiểu trả về interface{} cho phương thức Build()
// để cho phép trả về các đối tượng hình dạng khác nhau như Circle và Rectangle.
func GetShapeBuilder(builderType string) interface{} {
	switch builderType {
	case "circle":
		return &CircleBuilder{}
	case "rectangle":
		return &RectangleBuilder{}
	default:
		return nil
	}
	return nil
}

// Mẫu thiết kế cũ chỉ giúp ta trả về IShapeBuilder
// func GetShapeBuilderOld(builderType string) IShapeBuilder {
// 	if builderType == "Rectangle" {
// 		return &rectangleBuilder{}
// 	}
// 	if builderType == "Circle" {
// 		return &circleBuilder{}
// 	}
// 	return nil
// }
// Refer for more details: https://www.meisternote.com/app/note/nCcziE9SLxQw/getbuilder
