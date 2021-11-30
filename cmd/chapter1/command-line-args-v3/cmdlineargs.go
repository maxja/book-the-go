package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Thats a better way to do it
	fmt.Println(strings.Join(os.Args[1:], " "))
}
