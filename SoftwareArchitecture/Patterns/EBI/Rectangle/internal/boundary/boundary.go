package boundary

import (
	"fmt"

	"github.com/huavanthong/design-patterns/SoftwareArchitecture/Patterns/EBI/Rectangle/internal/entity"
)

type RectangleInput struct {
	Width  float64
	Height float64
}

type RectangleOutput struct {
	Area      float64
	Perimeter float64
}

type RectangleBoundary struct{}

func (b *RectangleBoundary) Calculate(input *RectangleInput) (*RectangleOutput, error) {
	if input.Width <= 0 || input.Height <= 0 {
		return nil, fmt.Errorf("width and height must be greater than 0")
	}

	rect := &entity.Rectangle{
		Width:  input.Width,
		Height: input.Height,
	}

	output := &RectangleOutput{
		Area:      rect.Width * rect.Height,
		Perimeter: 2 * (rect.Width + rect.Height),
	}

	return output, nil
}
