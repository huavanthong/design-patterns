package main

import (
	"fmt"
	"github.com/huavanthong/design-patterns/SoftwareArchitecture/Patterns/EBI/Rectangle/internal/boundary"
	"github.com/huavanthong/design-patterns/SoftwareArchitecture/Patterns/EBI/Rectangle/internal/interactor"
)

func main() {
	interactor := &interactor.RectangleInteractor{
		Boundary: &boundary.RectangleBoundary{},
	}

	input := &boundary.RectangleInput{
		Width:  10,
		Height: 20,
	}

	output, err := interactor.Calculate(input)s
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Area: %f\n", output.Area)
		fmt.Printf("Perimeter: %f\n", output.Perimeter)
	}
}
