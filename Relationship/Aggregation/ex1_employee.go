package main

import "fmt"

type Address struct {
	addressLine string
	city        string
	state       string
}

func NewAddress(addressLine string, city string, state string) Address {
	return Address{
		addressLine: addressLine,
		city:        city,
		state:       state,
	}
}

type Employee struct {
	id      int
	name    string
	address Address
}

/*
*****************************************************************************
	new employee by value
*****************************************************************************
*/
func NewEmployee(id int, name string, address Address) *Employee {
	e := Employee{
		id:      id,
		name:    name,
		address: address,
	}
	fmt.Printf("Check address after NewEmployee %p\n", &e)

	return &e
}

/*
*****************************************************************************
	new employee by pointer
*****************************************************************************
*/
func NewEmployeeByPointer(id int, name string, address Address) *Employee {

	e := new(Employee)
	e.id = id
	e.name = name
	e.address = address
	fmt.Printf("Check address after NewEmployeeByPointer %p\n", &e)

	return e
}

/*
*****************************************************************************
	main
*****************************************************************************
*/
func main() {
	address := NewAddress("Hau Giang", "HCM", "70000")
	fmt.Printf("Check address initialize %p\n", &address)

	/******************************************************************************
		create a object using constructors by value
	******************************************************************************/
	employee := NewEmployee(1, "Hua Van Thong", address)
	fmt.Printf("ID: %d ,\nName: %s, \n AddressLine: %s, Address of employee.address %p \n",
		employee.id,
		employee.name,
		employee.address.addressLine,
		&employee.address)

	/******************************************************************************
		create a object using constructors by pointer
	******************************************************************************/
	employee2 := NewEmployeeByPointer(1, "Hua Van Thong", address)
	fmt.Printf("ID: %d ,\nName: %s, \n AddressLine: %s, Address of employee.address %p\n",
		employee2.id,
		employee2.name,
		employee2.address.addressLine,
		&employee2.address)
}
