package interactor

import (
	"github.com/huavanthong/design-patterns/SoftwareArchitecture/Patterns/EBI/Rectangle/internal/boundary"
)

type RectangleInteractor struct {
	Boundary *boundary.RectangleBoundary
}

func (i *RectangleInteractor) Calculate(input *boundary.RectangleInput) (*boundary.RectangleOutput, error) {
	return i.Boundary.Calculate(input)
}
