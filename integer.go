package main

import (
	"fmt"
)

func main() {
	var x uint = 500
	var y uint = 4500
	var z int = 500
	var a int = -200
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)

	fmt.Printf("Type: %T, value: %v\n", z, z)
	fmt.Printf("Type: %T, value: %v\n", a, a)
}
