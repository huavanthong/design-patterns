package main

import (
	"fmt"

	"github.com/huavanthong/design-patterns/APP/boundary"
	bio "github.com/huavanthong/design-patterns/APP/boundary/io"

	"github.com/huavanthong/design-patterns/APP/builder"
	"github.com/huavanthong/design-patterns/APP/common"
	"github.com/huavanthong/design-patterns/APP/interactor"
	"github.com/huavanthong/design-patterns/APP/repository"
	"github.com/huavanthong/design-patterns/APP/validator"
)

func main() {

	// Initialize dependencies
	validator := validator.ShapeValidator{}

	rectangleRepository := repository.NewRectangleRepository()

	rectangleBuilder := builder.GetShapeBuilder("rectangle").(*builder.RectangleBuilder)

	// Set dimensions
	rectangleBuilder.SetDimensions(common.Dimensions{Width: 10, Height: 5, Radius: 5})

	// Set position
	rectangleBuilder.SetPosition(common.Position{X: 0, Y: 0})

	// Set color
	rectangleBuilder.SetColor(common.Color{R: 255, G: 0, B: 0})

	// Convert rectangleBuilder to entity rectangle
	rect := rectangleBuilder.Build()

	// Draw rectangle
	fmt.Println(rect.Draw())

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

	// Draw circle
	fmt.Println(circle.Draw())

	// Create a rectangle
	createInput := bio.CreateRectangleInput{
		ObjectName: "rectangle",
		OwnerName:  "thong",
		Dimensions: common.Dimensions{Width: 5, Height: 5, Radius: 0},
		Position:   common.Position{X: 1, Y: 1},
		Color:      common.Color{R: 1, B: 1, G: 1},
	}

	rectangleInteractor := interactor.NewRectangleInteractor(*rectangleRepository)

	// Initialize boundary instances
	rectangleBoundary := boundary.NewRectangleBoundary(rectangleInteractor, validator)

	createOutput, err := rectangleBoundary.CreateRectangle(createInput)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Raw info rectangle from entity: ", createOutput)

	//fmt.Println("Success: ", rectangleBoundary.SuccessRectangle())

}
