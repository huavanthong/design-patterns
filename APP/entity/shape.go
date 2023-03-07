package entity

type Shape interface {
	GetArea() float64
	GetPerimeter() float64
	Draw() string
}
