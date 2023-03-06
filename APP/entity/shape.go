package entity

import (
	"github.com/huavanthong/design-patterns/APP/common"
)

type shape struct {
	dimensions common.Dimensions
	position   common.Position
	color      common.Color
	borderSize int
}

type Shape interface {
	GetArea() float64
	GetPerimeter() float64
	Draw()
}
