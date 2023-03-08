package common

import (
	"github.com/google/uuid"
)

// GenerateUUID generates a unique identifier using UUID v4
func GenerateUUID() string {
	return uuid.NewString()
}

type Dimensions struct {
	Width  float64
	Height float64
	Radius float64
}

type Position struct {
	X float64
	Y float64
}

type Color struct {
	R uint8
	G uint8
	B uint8
}
