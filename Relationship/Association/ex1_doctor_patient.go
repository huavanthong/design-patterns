package main

import "fmt"

type Patient struct {
	name string
}

func (p Patient) getName() string {
	return p.name
}

type Doctor struct {
	name     string
	patients []Patient
}

func (d Doctor) getName() string {
	return d.name
}

func (d *Doctor) addPatient(p Patient) {
	d.patients = append(d.patients, p)
}

func main() {
	// Initialize value for structure
	var dave Patient = Patient{name: "Dave"}

	var bestsy Patient
	bestsy.name = "Betsy"

	frank := Patient{name: "Frank"}

	james := Doctor{name: "James"}
	scott := Doctor{name: "Scott"}

	james.addPatient(dave)

	scott.addPatient(bestsy)
	scott.addPatient(frank)

	fmt.Printf("Doctor: %s , Patients: %s\n", james.name, james.patients[0])

	fmt.Printf("Doctor: %s , Patients: %s, %s\n", scott.name, scott.patients[0], scott.patients[1])

}
