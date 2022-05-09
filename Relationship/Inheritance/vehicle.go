package main

import "fmt"

//Define a behavior for vehicles
type VehicleBehavior interface {
	PrintType()
}

type Car struct {
}

func (Car) PrintType() {
	fmt.Println("This is my car")
}

type Bus struct {
}

func (Bus) PrintType() {
	fmt.Println("This is my bus")
}

type Truck struct {
}

func (Truck) PrintType() {
	fmt.Println("This is my truck")
}

func main() {

	var car Car
	car.PrintType()

	var bus Bus
	bus.PrintType()

	var truck Truck
	truck.PrintType()
}
