package interactor

import (
	"github.com/huavanthong/design-patterns/APP/entity"
	"github.com/huavanthong/design-patterns/APP/repository"
)

type ShapeUseCase struct {
	rectangleRepository entity.RectangleRepository
}
type shapeInteractor struct {
	rectangleRepository repository.RectangleRepository
}

func NewMyUseCase(rr entity.RectangleRepository) *ShapeUseCase {
	return &ShapeUseCase{rectangleRepository: rr}
}

func (si *shapeInteractor) FindShapesByColor(color string) ([]*entity.Shape, error) {
	shapes, err := si.shapeRepository.GetAllShapes()
	if err != nil {
		return nil, err
	}

	var matchedShapes []*entity.Shape
	for _, shape := range shapes {
		if shape.Color == color {
			matchedShapes = append(matchedShapes, shape)
		}
	}

	return matchedShapes, nil
}

func (i *shapeInteractor) CalculateTotalArea(shapes []*entity.Shape) float64 {
	var totalArea float64
	for _, shape := range shapes {
		totalArea += shape.CalculateArea()
	}

	return totalArea
}

func (i *shapeInteractor) FindShapeWithMaxArea(shapes []*entity.Shape) *entity.Shape {
	var maxAreaShape *entity.Shape
	var maxArea float64
	for _, shape := range shapes {
		if area := shape.CalculateArea(); area > maxArea {
			maxArea = area
			maxAreaShape = shape
		}
	}

	return maxAreaShape
}
