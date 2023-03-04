package boundary

import (
	"fmt"

	"github.com/huavanthong/design-patterns/SoftwareArchitecture/Patterns/EBI/Circle/internal/entity"
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
type RectangleBoundary struct{}

func (b *RectangleBoundary) Calculate(input *CircleInput) (*CircleOutput, error) {
	if input.Radius <= 0 {
		return nil, fmt.Errorf("Invalid input")
	}

	rect := &entity.Circle{
		Radius: input.Radius,
	}

	output := &CircleOutput{
		Area:      rect.Width * rect.Height,
		Perimeter: 2 * (rect.Width + rect.Height),
	}

	return output, nil
}
