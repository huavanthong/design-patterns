package main

import (
	"fmt"
	"math"
)

type Dimensions struct {
	Width  float64
	Height float64
	Radius float64
}

type Color struct {
	R uint8
	G uint8
	B uint8
}
type Position struct {
	X float64
	Y float64
}

type ShapeBuilder interface {
	SetDimensions(dimensions Dimensions) ShapeBuilder
	SetColor(color Color) ShapeBuilder
	SetPosition(position Position) ShapeBuilder
	Build() Shape
}

type CircleBuilder interface {
	SetDimensions(dimensions Dimensions) CircleBuilder
	SetColor(color Color) CircleBuilder
	SetPosition(position Position) CircleBuilder
	Build() Circle
}

type RectangleBuilder interface {
	SetDimensions(dimensions Dimensions) RectangleBuilder
	SetColor(color Color) RectangleBuilder
	SetPosition(position Position) RectangleBuilder
	Build() Rectangle
}

type Shape interface {
	GetArea() float64
	GetPerimeter() float64
	Draw() string
}

type Circle struct {
	Dimensions Dimensions
	Color      Color
	Position   Position
}

func (c Circle) GetArea() float64 {
	return math.Pi * c.Dimensions.Radius * c.Dimensions.Radius
}

func (c Circle) GetPerimeter() float64 {
	return 2 * math.Pi * c.Dimensions.Radius
}

func (c Circle) Draw() string {
	return fmt.Sprintf("Drawing Circle with radius %v at position (%v, %v)", c.Dimensions.Radius, c.Position.X, c.Position.Y)
}

type Rectangle struct {
	Dimensions Dimensions
	Color      Color
	Position   Position
}

func (r Rectangle) GetArea() float64 {
	return r.Dimensions.Width * r.Dimensions.Height
}

func (r Rectangle) GetPerimeter() float64 {
	return 2 * (r.Dimensions.Width + r.Dimensions.Height)
}

func (r Rectangle) Draw() string {
	return fmt.Sprintf("Drawing Rectangle with width %v, height %v at position (%v, %v)", r.Dimensions.Width, r.Dimensions.Height, r.Position.X, r.Position.Y)
}

type circleBuilder struct {
	dimensions Dimensions
	color      Color
	position   Position
}

func (c *circleBuilder) SetDimensions(dimensions Dimensions) CircleBuilder {
	c.dimensions = dimensions
	return c
}

func (c *circleBuilder) SetColor(color Color) CircleBuilder {
	c.color = color
	return c
}

func (c *circleBuilder) SetPosition(position Position) CircleBuilder {
	c.position = position
	return c
}

func (c *circleBuilder) Build() Circle {
	circle := Circle{
		Dimensions: c.dimensions,
		Color:      c.color,
		Position:   c.position,
	}
	return circle
}

type rectangleBuilder struct {
	dimensions Dimensions
	color      Color
	position   Position
}

func (r *rectangleBuilder) SetDimensions(dimensions Dimensions) RectangleBuilder {
	r.dimensions = dimensions
	return r
}

func (r *rectangleBuilder) SetColor(color Color) RectangleBuilder {
	r.color = color
	return r
}

func (r *rectangleBuilder) SetPosition(position Position) RectangleBuilder {
	r.position = position
	return r
}

func (r *rectangleBuilder) Build() Rectangle {
	rectangle := Rectangle{
		Dimensions: r.dimensions,
		Color:      r.color,
		Position:   r.position,
	}
	return rectangle
}

func GetBuilder(shapeType string) interface{} {
	switch shapeType {
	case "circle":
		return &circleBuilder{}
	case "rectangle":
		return &rectangleBuilder{}
	default:
		return nil
	}
}
func main() {
	circleBuilder := GetBuilder("circle").(*circleBuilder)
	circle := circleBuilder.SetDimensions(Dimensions{Radius: 5}).
		SetColor(Color{R: 255, G: 0, B: 0}).
		SetPosition(Position{X: 10, Y: 20}).
		Build()
	fmt.Println(circle)

	rectangleBuilder := GetBuilder("rectangle").(*rectangleBuilder)
	rectangle := rectangleBuilder.SetDimensions(Dimensions{Width: 10, Height: 20}).
		SetColor(Color{R: 0, G: 255, B: 0}).
		SetPosition(Position{X: 30, Y: 40}).
		Build()
	fmt.Println(rectangle)
}
