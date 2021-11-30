package main

import (
	"fmt"
	"os"
)

func main() {
	// For the variable keyword var can be used
	var s, sep string
	// For is the one and only loop operator on Go
	for i := 1; i < len(os.Args); i++ {
		// Depending on a type operators can be overloaded
		// For strings it works as concatenation
		// This is an awful way to extend string
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
