package boundary

import (
	"fmt"

	bio "github.com/huavanthong/design-patterns/APP/boundary/io"
	"github.com/huavanthong/design-patterns/APP/builder"
	"github.com/huavanthong/design-patterns/APP/common"
)

// circle_boundary.go
type CircleBoundary struct{}

func (b *CircleBoundary) Calculate(input *bio.CircleInput) (*bio.CircleOutput, error) {
	if input.Radius <= 0 {
		return nil, fmt.Errorf("Invalid input")
	}

	circleBuilder := builder.GetShapeBuilder("circle").(*builder.CircleBuilder)

	circleBuilder.SetDimensions(common.Dimensions{Width: 0, Height: 0, Radius: input.Radius})
	circle := circleBuilder.Build()

	output := &bio.CircleOutput{
		Area:      circle.GetArea(),
		Perimeter: circle.GetPerimeter(),
	}

	return output, nil
}
