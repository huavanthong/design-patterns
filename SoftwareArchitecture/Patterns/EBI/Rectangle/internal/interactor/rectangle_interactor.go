package interactor

import (
	"github.com/huavanthong/design-patterns/SoftwareArchitecture/Patterns/EBI/Circle/internal/boundary"
	"github.com/huavanthong/design-patterns/SoftwareArchitecture/Patterns/EBI/Circle/internal/builder"
)

type RectangleInteractor struct {
	input  boundary.RectangleInput
	output boundary.RectangleOutput
}

func NewRectangleInteractor(input boundary.RectangleInput, output boundary.RectangleOutput) *RectangleInteractor {
	return &RectangleInteractor{
		input:  input,
		output: output,
	}
}

func (ri *RectangleInteractor) CreateRectangle(builder builder.ShapeBuilder) error {
	// Lấy thông tin input từ boundary
	width, height, err := ri.input.GetRectangleInput()
	if err != nil {
		return err
	}

	// Tạo dimensions cho Rectangle bằng builder pattern
	dimensions := builder.
		SetWidth(width).
		SetHeight(height).
		BuildDimensions()

	// Tạo Rectangle từ dimensions
	rectangle := builder.
		SetDimensions(dimensions).
		Build()

	// Gửi thông tin Rectangle tới output boundary
	err = ri.output.OutputRectangle(rectangle)
	if err != nil {
		return err
	}

	return nil
}
