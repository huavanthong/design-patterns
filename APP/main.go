package main

import (
	"fmt"
)

func main() {
	// Create rectangle builder
	rectangleBuilder := NewRectangleBuilder()

	// Set dimensions
	rectangleBuilder.SetDimensions(Dimensions{Width: 10, Height: 5, Radius: 5})

	// Set position
	rectangleBuilder.SetPosition(Position{X: 0, Y: 0})

	// Set color
	rectangleBuilder.SetColor(Color{R: 255, G: 0, B: 0})

	// Build rectangle
	rectangle := rectangleBuilder.Build()

	// Create circle builder
	circleBuilder := NewCircleBuilder()

	// Set dimensions
	circleBuilder.SetRadius(7)

	// Set position
	circleBuilder.SetPosition(Position{X: 2, Y: 2})

	// Set color
	circleBuilder.SetColor(Color{R: 0, G: 255, B: 0})

	// Build circle
	circle := circleBuilder.Build()

	// Create rectangle interactor
	rectangleInteractor := NewRectangleInteractor(rectangle)

	// Create circle interactor
	circleInteractor := NewCircleInteractor(circle)

	// Print rectangle info
	fmt.Println(rectangleInteractor.GetInfo())

	// Print circle info
	fmt.Println(circleInteractor.GetInfo())
}
