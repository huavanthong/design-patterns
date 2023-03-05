package main

import (
	"fmt"

	"github.com/huavanthong/design-patterns/APP/builder"
	"github.com/huavanthong/design-patterns/APP/interactor"

	"github.com/huavanthong/design-patterns/APP/common"
)

func main() {
	// Create rectangle builder
	rectangleBuilder := builder.NewRectangleBuilder()

	// Set dimensions
	rectangleBuilder.SetDimensions(common.Dimensions{Width: 10, Height: 5, Radius: 5})

	// Set position
	rectangleBuilder.SetPosition(common.Position{X: 0, Y: 0})

	// Set color
	rectangleBuilder.SetColor(common.Color{R: 255, G: 0, B: 0})

	// Build rectangle
	rectangle := rectangleBuilder.Build()

	// Create circle builder
	circleBuilder := builder.NewCircleBuilder()

	// Set dimensions
	circleBuilder.SetRadius(7)

	// Set position
	circleBuilder.SetPosition(common.Position{X: 2, Y: 2})

	// Set color
	circleBuilder.SetColor(common.Color{R: 0, G: 255, B: 0})

	// Build circle
	circle := circleBuilder.Build()

	// Create rectangle interactor
	rectangleInteractor := interactor.NewRectangleInteractor(rectangle)

	// Create circle interactor
	circleInteractor := interactor.NewCircleInteractor(circle)

	// Print rectangle info
	fmt.Println(rectangleInteractor.GetInfo())

	// Print circle info
	fmt.Println(circleInteractor.GetInfo())
}
