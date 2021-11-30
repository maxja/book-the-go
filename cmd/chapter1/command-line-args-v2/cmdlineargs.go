package main

import (
	"fmt"
	"os"
)

func main() {
	// Declaration can be done with initialisation
	// and inside a function it can be merged into := form
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
