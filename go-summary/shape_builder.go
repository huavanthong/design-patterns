package main

import "fmt"

type ShapeBuilder interface {
	SetX(int)
	SetY(int)
	SetFillColor(string)
}

type Circle struct {
	X         int
	Y         int
	FillColor string
	Radius    int
}

type CircleBuilder struct {
	Circle
}

func (cb *CircleBuilder) SetX(x int) {
	cb.X = x
}

func (cb *CircleBuilder) SetY(y int) {
	cb.Y = y
}

func (cb *CircleBuilder) SetFillColor(fillColor string) {
	cb.FillColor = fillColor
}

func (cb *CircleBuilder) SetRadius(radius int) {
	cb.Radius = radius
}

func (cb *CircleBuilder) Build() *Circle {
	return &cb.Circle
}

type Rectangle struct {
	X1        int
	Y1        int
	X2        int
	Y2        int
	FillColor string
}

type RectangleBuilder struct {
	Rectangle
}

func (rb *RectangleBuilder) SetX1(x1 int) {
	rb.X1 = x1
}

func (rb *RectangleBuilder) SetY1(y1 int) {
	rb.Y1 = y1
}

func (rb *RectangleBuilder) SetX2(x2 int) {
	rb.X2 = x2
}

func (rb *RectangleBuilder) SetY2(y2 int) {
	rb.Y2 = y2
}

func (rb *RectangleBuilder) SetFillColor(fillColor string) {
	rb.FillColor = fillColor
}

func (rb *RectangleBuilder) Build() *Rectangle {
	return &rb.Rectangle
}

func main() {
	circleBuilder := &CircleBuilder{}
	circleBuilder.SetX(10)
	circleBuilder.SetY(20)
	circleBuilder.SetRadius(30)
	circleBuilder.SetFillColor("red")
	circle := circleBuilder.Build()
	fmt.Println(circle)

	rectangleBuilder := &RectangleBuilder{}
	rectangleBuilder.SetX1(0)
	rectangleBuilder.SetY1(0)
	rectangleBuilder.SetX2(100)
	rectangleBuilder.SetY2(50)
	rectangleBuilder.SetFillColor("blue")
	rectangle := rectangleBuilder.Build()
	fmt.Println(rectangle)
}
