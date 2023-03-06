package boundary

import (
	"fmt"

	"github.com/huavanthong/design-patterns/APP/builder"
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

	circle := builder.NewCircleBuilder().
		SetDimensions(input.Radius).
		Build()

	output := &CircleOutput{
		Area:      circle.GetArea(),
		Perimeter: circle.GetPerimeter(),
	}

	return output, nil
}
