/********************************************************************************************

Function name: Test suite for Shape

Program name: shape_test.go

Description: Implement test case for Shape

Creator :Van Thong

Creation date :202/04/15

Change log	:

********************************************************************************************/
package entity

import (
	"fmt"
	"testing"
)

// Test inheritance for method GetArea for all shapes
func TestShape_GetArea(t *testing.T) {
	// Set up test data
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Square{Size: 2},
		Circle{Radius: 2},
	}

	// Define expected results
	expected := []float64{12, 4, 12.566370614359172}

	for i, shape := range shapes {
		result := shape.GetArea()
		if result != expected[i] {
			t.Errorf("Shape %d: expected area %f but got %f", i+1, expected[i], result)
		}
	}
}

// Test inheritance for method GetPerimeter for all shapes
func TestShape_GetPerimeter(t *testing.T) {
	// Set up test data
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Square{Size: 2},
		Circle{Radius: 2},
	}

	// Define expected results
	expected := []float64{14, 8, 12.566370614359172}

	// Iterate over shapes and test GetPerimeter method
	for i, shape := range shapes {
		result := shape.GetPerimeter()
		if result != expected[i] {
			t.Errorf("Shape %d: expected perimeter %f but got %f", i+1, expected[i], result)
		}
	}
}

// Test inheritance for method Draw for all shapes
func TestShape_Draw(t *testing.T) {
	// Set up test data
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Square{Size: 2},
		Circle{Radius: 2},
	}

	for i, shape := range shapes {
		fmt.Printf("Shape %d: Show draw info: %v\n", i+1, shape.Draw())
	}
}
