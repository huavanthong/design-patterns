package main

import (
	"errors"
	"fmt"
	"time"
)

/*

type error interface {
    Error() string
}


Refer: https://go.dev/tour/methods/19
*/

// Step 1: We define my error with my expection. Error retrieve 2 info than one string info as old error
type MyError struct {
	When time.Time
	What string
}

// Step 2: we know that error is interface, so we can override them following our struct error info
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// Step 3: implement method to test our error.
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {

	fmt.Println("Check old error: ", errors.New("temp"))

	if err := run(); err != nil {
		fmt.Println(err)
	}
}
