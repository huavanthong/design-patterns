package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		go getInstance()
	}

	go removeInstance()
	fmt.Println("Remove success")
	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
