package repository

import (
	"github.com/huavanthong/design-patterns/APP/common"
	"github.com/huavanthong/design-patterns/APP/entity"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircleRepository_Save(t *testing.T) {
	// Set up test data
	id := "circle-1"
	dimension := common.Dimensions{Width: 10, Height: 5, Radius: 0}
	circle := &entity.Circle{
		ID:         id,
		ObjectName: "Circle 1",
		OwnerName:  "Van Thong",
		Radius:     dimension.Radius,
		Color:      common.Color{R: 255, G: 255, B: 255},
		Position:   common.Position{X: 0, Y: 0},
		BorderSize: 1,
	}

	repo := NewInMemoryCircleRepository()

	// Test case 1: Valid ID
	err := repo.Save(circle)
	assert.NoError(t, err)
}
