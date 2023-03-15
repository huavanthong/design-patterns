package repository

import (
	"github.com/huavanthong/design-patterns/APP/common"
	"github.com/huavanthong/design-patterns/APP/entity"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShapeRepository_Save(t *testing.T) {
	// Set up test data
	id := "rect-1"
	dimension := common.Dimensions{Width: 10, Height: 5, Radius: 0}
	rect := &entity.Rectangle{
		ID:         id,
		ObjectName: "Rectangle 1",
		OwnerName:  "Van Thong",
		Width:      dimension.Width,
		Height:     dimension.Height,
		Color:      common.Color{R: 255, G: 255, B: 255},
		Position:   common.Position{X: 0, Y: 0},
		BorderSize: 1,
	}

	repo := NewRectangleRepository()

	// Test case 1: Valid ID
	err := repo.Save(rect)
	assert.NoError(t, err)
}
