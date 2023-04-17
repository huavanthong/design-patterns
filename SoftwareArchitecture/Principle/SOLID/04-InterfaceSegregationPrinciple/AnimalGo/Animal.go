package main

import "fmt"

type Animal struct {
	ID   int
	Name string
}

type Flyable interface {
	fly()
}

type Runable interface {
	run()
}

type Barkable interface {
	bark()
}

type Cat struct {
	Animal
}

func (c Cat) run() {
	fmt.Println("Cat is running")
}

type Dog struct {
	Animal
}

func (d Dog) run() {
	fmt.Println("Cat is running")
}

func (d Dog) bark() {
	fmt.Println("Cat is barkin")
}

type Bird struct {
	Animal
}

func (b Bird) bark() {
	fmt.Println("Bird is leo lo leo lo")
}

func (b Bird) fly() {
	fmt.Println("Bird is flying")
}

func main() {
	cat := Cat{Animal: Animal{
		ID:   1,
		Name: "cat",
	}}
	cat.run()
	fmt.Println(cat)

}
