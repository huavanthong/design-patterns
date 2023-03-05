package main

import (
	"fmt"
)

// Ta có thể đặt file này ở file khác chẳng hạn circle_input.go
type CircleInput struct {
	Radius float64
}

// Ta có thể đặt file này ở file khác chẳng hạn circle_output.go
type CircleOutput struct {
	Area      float64
	Perimeter float64
}

// circle_boundary.go
type CircleBoundary struct{}

func (b *CircleBoundary) Calculate(input *CircleInput) (*CircleOutput, error) {
	if input.Radius <= 0 {
		return nil, fmt.Errorf("Invalid input")
	}

	rect := &Circle{
		radius: input.Radius,
	}

	output := &CircleOutput{
		Area:      rect.radius * rect.radius,
		Perimeter: 2 * (rect.radius + rect.radius),
	}

	return output, nil
}
