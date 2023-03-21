package repository

import (
	"github.com/huavanthong/design-patterns/APP/entity"
)

// ShapeRepository defines the interface for storing and retrieving shapes.
/*
Experience 8: Sử dụng nguyên tắc Dependency Inversion Principle trong SOLID
	Chúng ta triển khai "ShapeRepository" và hai implementation của nó là "RectangleRepository" và "CircleRepository"
	bằng cách sử dụng interface trong ngôn ngữ Golang.

Qua đó, chúng ta đứng từ lớp cao nhất. Và cung cấp các phương thức cơ bản nhất cho các implementation còn lại.
*/
type ShapeRepository interface {
	// Experience 1: Experience in using pointer and variable in Go
	Create(shape *entity.Shape)
	GetByID(id string) ([]entity.Shape, error)
	Update(shape *entity.Shape) error
	DeleteByID(ID string) error
}
