package boundary

import (
	"fmt"

	"github.com/huavanthong/design-patterns/APP/builder"
	"github.com/huavanthong/design-patterns/APP/common"
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

	circleBuilder := builder.GetShapeBuilder("circle").(*builder.CircleBuilder)

	circleBuilder.SetDimensions(common.Dimensions{Width: 0, Height: 0, Radius: input.Radius})
	circle := circleBuilder.Build()

	output := &CircleOutput{
		Area:      circle.GetArea(),
		Perimeter: circle.GetPerimeter(),
	}

	return output, nil
}
