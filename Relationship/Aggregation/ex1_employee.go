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

func NewEmployee(id int, name string, address Address) Employee {
	return Employee{
		id:      id,
		name:    name,
		address: address,
	}
}

func main() {
	address := NewAddress("Hau Giang", "HCM", "70000")
	employee := NewEmployee(1, "Hua Van Thong", address)
	fmt.Printf("ID: %d ,\nName: %s, \n AddressLine: %s, \n",
		employee.id,
		employee.name,
		employee.address.addressLine)
}
