package main

import (
	"fmt"

	"github.com/huavanthong/design-patterns/APP/builder"
	"github.com/huavanthong/design-patterns/APP/common"
	"github.com/huavanthong/design-patterns/APP/interactor"
	"github.com/huavanthong/design-patterns/APP/validator"
)

func main() {
	// Create rectangle builder
	rectangleBuilder := builder.GetShapeBuilder("rectangle").(*builder.RectangleBuilder)

	// Set dimensions
	rectangleBuilder.SetDimensions(common.Dimensions{Width: 10, Height: 5, Radius: 5})

	// Set position
	rectangleBuilder.SetPosition(common.Position{X: 0, Y: 0})

	// Set color
	rectangleBuilder.SetColor(common.Color{R: 255, G: 0, B: 0})

	// Build rectangle
	rectangle := rectangleBuilder.Build()

	// Create circle builder
	circleBuilder := builder.GetShapeBuilder("circle").(*builder.CircleBuilder)

	// Set dimensions
	circleBuilder.SetDimensions(common.Dimensions{Width: 10, Height: 5, Radius: 5})

	// Set position
	circleBuilder.SetPosition(common.Position{X: 2, Y: 2})

	// Set color
	circleBuilder.SetColor(common.Color{R: 0, G: 255, B: 0})

	// Build circle
	circle := circleBuilder.Build()

	validator := validator.NewShapeValidator()

	// Create rectangle interactor
	rectangleInteractor := interactor.NewRectangleInteractor(rectangle, validator)

	// // Create circle interactor
	// circleInteractor := interactor.NewCircleInteractor(circle)

	// Print rectangle info
	fmt.Println(rectangleInteractor.GetInfo())

	// Print circle info
	fmt.Println(circle)
}
