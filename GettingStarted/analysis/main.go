package main

import (
	h "analysis/human"
	v "analysis/vehicle"

	"fmt"
)

func main() {
	// Step 3: Get interface from vehicle
	var bmw v.Vehicle

	// Step 4: We can use a global variable for it
	bmw = v.Car("World Top Brand")

	var labour h.Human
	labour = h.Man("Software Developer")

	// Step 5: And we also use all methods implement from Interface
	for i, j := range bmw.Structure() {
		fmt.Printf("%-15s <=====> %15s\n", j, labour.Structure()[i])
	}

	printStructure(bmw)
}

func printStructure(bmw v.Vehicle) {
	for _, j := range bmw.Structure() {
		fmt.Printf("%-15s\n", j)
	}
}
