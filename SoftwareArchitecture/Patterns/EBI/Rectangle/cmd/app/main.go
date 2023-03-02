package main

import (
	"fmt"

	"github.com/example/project/boundary"
	"github.com/example/project/interactor"
)

func main() {
	interactor := &interactor.RectangleInteractor{
		Boundary: &boundary.RectangleBoundary{},
	}

	input := &boundary.RectangleInput{
		Width:  10,
		Height: 20,
	}

	output, err := interactor.Calculate(input)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Area: %f\n", output.Area)
		fmt.Printf("Perimeter: %f\n", output.Perimeter)
	}
}
