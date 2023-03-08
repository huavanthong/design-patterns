package main

import (
	"fmt"

	"github.com/huavanthong/design-patterns/APP/boundary"
	"github.com/huavanthong/design-patterns/APP/builder"
	"github.com/huavanthong/design-patterns/APP/common"
	"github.com/huavanthong/design-patterns/APP/interactor"
	"github.com/huavanthong/design-patterns/APP/repository"
	"github.com/huavanthong/design-patterns/APP/validator"
)

func main() {
	/****************************** Use Case 1 *******************************/
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

	/****************************** Use Case 2 *******************************/
	// Create repository and boundary instances
	rectRepo := repository.NewInMemoryRectangleRepository()
	circleRepo := repository.NewInMemoryCircleRepository()
	rectBoundary := boundary.NewRectangleBoundary(rectRepo)
	circleBoundary := boundary.NewCircleBoundary(circleRepo)

	// Create validator instance
	shapeValidator := validator.NewShapeValidator()

	// Create interactor instances
	rectInteractor := interactor.NewRectangleInteractor(rectBoundary, shapeValidator)
	circleInteractor := interactor.NewCircleInteractor(circleBoundary, shapeValidator)

	// Create rectangle
	rectInput := boundary.CreateRectangleInput{
		Dimensions: boundary.Dimensions{
			Width:  10,
			Height: 20,
		},
	}
	rect, err := rectInteractor.Create(rectInput)
	if err != nil {
		fmt.Println("Error creating rectangle:", err)
		return
	}
	fmt.Println("Created rectangle:", rect)

	// Get rectangle by ID
	getRectInput := boundary.GetRectangleInput{
		ID: rect.ID,
	}
	rect, err = rectInteractor.Get(getRectInput)
	if err != nil {
		fmt.Println("Error getting rectangle:", err)
		return
	}
	fmt.Println("Got rectangle by ID:", rect)

	// Update rectangle
	updateRectInput := boundary.UpdateRectangleInput{
		ID: rect.ID,
		Dimensions: boundary.Dimensions{
			Width:  30,
			Height: 40,
		},
	}
	rect, err = rectInteractor.Update(updateRectInput)
	if err != nil {
		fmt.Println("Error updating rectangle:", err)
		return
	}
	fmt.Println("Updated rectangle:", rect)

	// Delete rectangle by ID
	deleteRectInput := boundary.DeleteRectangleInput{
		ID: rect.ID,
	}
	err = rectInteractor.Delete(deleteRectInput)
	if err != nil {
		fmt.Println("Error deleting rectangle:", err)
		return
	}
	fmt.Println("Deleted rectangle by ID:", rect.ID)

	// Create circle
	circleInput := boundary.CreateCircleInput{
		Dimensions: boundary.Dimensions{
			Radius: 5,
		},
	}
	circle, err := circleInteractor.Create(circleInput)
	if err != nil {
		fmt.Println("Error creating circle:", err)
		return
	}
	fmt.Println("Created circle:", circle)
}
