package repository

import (
	"github.com/huavanthong/design-patterns/APP/common"
	"github.com/huavanthong/design-patterns/APP/entity"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRectangleRepository_Save(t *testing.T) {
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

func TestRectangleRepository_FindByID(t *testing.T) {
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

	// Pre test
	repo.Save(rect)

	// Test case 1: Valid ID
	result, err := repo.FindByID(id)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, rect, result)

	// Test case 2: Invalid ID
	result, err = repo.FindByID("Invalid-ID")
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestRectangleRepository_FindAll(t *testing.T) {
	// Set up test data
	id1 := "rect-1"
	dimension1 := common.Dimensions{Width: 10, Height: 5, Radius: 0}
	rect1 := &entity.Rectangle{
		ID:         id1,
		ObjectName: "Rectangle 1",
		OwnerName:  "Van Thong",
		Width:      dimension1.Width,
		Height:     dimension1.Height,
		Color:      common.Color{R: 255, G: 255, B: 255},
		Position:   common.Position{X: 0, Y: 0},
		BorderSize: 1,
	}

	id2 := "rect-2"
	dimension2 := common.Dimensions{Width: 2, Height: 2, Radius: 0}
	rect2 := &entity.Rectangle{
		ID:         id2,
		ObjectName: "Rectangle 2",
		OwnerName:  "Van Thong",
		Width:      dimension2.Width,
		Height:     dimension2.Height,
		Color:      common.Color{R: 255, G: 255, B: 255},
		Position:   common.Position{X: 0, Y: 0},
		BorderSize: 1,
	}

	repo := NewRectangleRepository()

	// Pre test
	repo.Save(rect1)
	repo.Save(rect2)

	// Test case 1: Valid ID
	var rectangles []entity.Rectangle
	rectangles = append(rectangles, *rect1)
	rectangles = append(rectangles, *rect2)

	result, err := repo.FindAll()
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, rectangles, result)

}

func TestRectangleRepository_Update(t *testing.T) {
	// Set up test data
	id1 := "rect-1"
	dimension1 := common.Dimensions{Width: 10, Height: 5, Radius: 0}
	rect1 := &entity.Rectangle{
		ID:         id1,
		ObjectName: "Rectangle 1",
		OwnerName:  "Van Thong",
		Width:      dimension1.Width,
		Height:     dimension1.Height,
		Color:      common.Color{R: 255, G: 255, B: 255},
		Position:   common.Position{X: 0, Y: 0},
		BorderSize: 1,
	}

	id2 := "rect-1"
	dimension2 := common.Dimensions{Width: 2, Height: 2, Radius: 0}
	rect2 := &entity.Rectangle{
		ID:         id2,
		ObjectName: "Rectangle 2",
		OwnerName:  "Van Thong 2",
		Width:      dimension2.Width,
		Height:     dimension2.Height,
		Color:      common.Color{R: 255, G: 255, B: 255},
		Position:   common.Position{X: 0, Y: 0},
		BorderSize: 1,
	}

	repo := NewRectangleRepository()

	// Pre test
	repo.Save(rect1)

	// Test case 1: Update rect1 to rect2
	err := repo.Update(rect2)
	assert.NoError(t, err)

	result, err2 := repo.FindByID(id2)
	assert.NoError(t, err2)
	assert.NotNil(t, result)
	assert.Equal(t, rect2, result)

}

func TestRectangleRepository_DeleteByID(t *testing.T) {
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

	// Pre test
	repo.Save(rect)

	// Test case 1: Valid ID
	err := repo.DeleteByID(id)
	assert.NoError(t, err)
}
