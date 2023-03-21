package main

type Shape struct {
	ID   string
	Type string
}

type ShapeRepository interface {
	Create(shape *Shape)
	GetByID(id string) (*Shape, error)
	Update(shape *Shape) error
	DeleteByID(ID string) error
}

type ShapeRepositoryImpl struct {
}

func (c *ShapeRepositoryImpl) Create(shape *Shape) error {
	// implementation for creating circle
	panic(nil)
}

func (c *ShapeRepositoryImpl) GetByID(id string) (*Shape, error) {
	// implementation for getting circle by ID
	panic(nil)
}

func (c *ShapeRepositoryImpl) Update(shape *Shape) error {
	// implementation for updating circle
	panic(nil)
}

func (c *ShapeRepositoryImpl) DeleteByID(id string) error {
	// implementation for deleting circle by ID
	panic(nil)
}

type Circle struct {
	ID     string
	Radius float64
}

type CircleRepository interface {
	Create(circle *Circle) error
	GetByID(id string) (*Circle, error)
	Update(circle *Circle) error
	DeleteByID(id string) error
}

type CircleRepositoryImpl struct {
	ShapeRepository
}

func (c *CircleRepositoryImpl) Create(circle *Circle) error {
	// implementation for creating circle
	panic(nil)
}

func (c *CircleRepositoryImpl) GetByID(id string) (*Circle, error) {
	// implementation for getting circle by ID
	panic(nil)
}

func (c *CircleRepositoryImpl) Update(circle *Circle) error {
	// implementation for updating circle
	panic(nil)
}

func (c *CircleRepositoryImpl) DeleteByID(id string) error {
	// implementation for deleting circle by ID
	panic(nil)
}

func main() {
	circleRepo := &CircleRepositoryImpl{}
	circle := Circle{}
	circleRepo.Create(&circle) // create a circle using CircleRepository's Create method

	circleRepo2 := &CircleRepositoryImpl{}
	shape := Shape{}
	circleRepo2.Create(&shape) // create a shape using ShapeRepository's Create method
}
