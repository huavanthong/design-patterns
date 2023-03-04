package main

type RectangleInteractor struct {
	input  RectangleInput
	output RectangleOutput
}

func NewRectangleInteractor(input RectangleInput, output RectangleOutput) *RectangleInteractor {
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
