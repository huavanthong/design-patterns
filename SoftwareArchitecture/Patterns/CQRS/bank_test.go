package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	old := os.Stdout // backup old stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = old // restore old stdout
	fmt.Printf(string(out))

	// Run setup code here

	// Run all tests
	code := m.Run()

	// Run teardown code here

	os.Exit(code)
}

/****************************** Use Case 1 *******************************/
func TestUseCase_01(t *testing.T) {
	fmt.Println("No test case")
}
