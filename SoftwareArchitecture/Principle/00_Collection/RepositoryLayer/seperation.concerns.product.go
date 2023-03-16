package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
}

type ProductRepository interface {
	Save(product *Product) error
	GetByID(id int) (*Product, error)
	Update(product *Product) error
	Delete(id int) error
}

type ProductService interface {
	GetProduct(id int) (*Product, error)
	CreateProduct(product *Product) error
	UpdateProduct(product *Product) error
	DeleteProduct(id int) error
}

type productLogic struct {
	repository ProductRepository
}

func (p *productLogic) GetProduct(id int) (*Product, error) {
	return p.repository.GetByID(id)
}

func (p *productLogic) CreateProduct(product *Product) error {
	return p.repository.Save(product)
}

func (p *productLogic) UpdateProduct(product *Product) error {
	return p.repository.Update(product)
}

func (p *productLogic) DeleteProduct(id int) error {
	return p.repository.Delete(id)
}

func NewProductService(repository ProductRepository) ProductService {
	return &productLogic{repository: repository}
}

func main() {
	// Create productLogic instance
	productLogic := productLogic{}

	// Create product instance
	product := Product{
		ID:          1,
		Name:        "Tivi",
		Description: "LG-AUT",
		Price:       18.5,
	}

	// Call service
	productLogic.CreateProduct(&product)

	result, err := productLogic.GetProduct(1)
	if err != nil {
		fmt.Printf("Info: %v", result)
	}
}
