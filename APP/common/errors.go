package common

import "errors"

var (
	// ErrInvalidWidth is the error for invalid rectangle width.
	ErrInvalidWidth = errors.New("Invalid width")
	// ErrInvalidHeight is the error for invalid rectangle height.
	ErrInvalidHeight = errors.New("Invalid height")
	// ErrInvalidRadius is the error for invalid circle radius.
	ErrInvalidRadius = errors.New("Invalid Radius")
	// Other errors ...
)
