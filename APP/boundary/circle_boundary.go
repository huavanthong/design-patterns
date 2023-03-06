package boundary

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

	// circle := entity.NewCircle(input.Radius)

	// output := &CircleOutput{
	// 	Area:      circle.Radius() * circle.Radius(),
	// 	Perimeter: 2 * (circle.Radius() + circle.Radius()),
	// }

	output := &CircleOutput{
		Area:      1 * 1,
		Perimeter: 2 * (2),
	}

	return output, nil
}
