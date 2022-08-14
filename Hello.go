package main

import (
	"fmt"
)

// Global Scope
var a int = 10

func add(x, y int) (int, int) {
	var out = x + y
	var out1 = x - y
	return out, out1

}

func main() {
	// FOR LOOP
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	// result1, result2 := add(1, 2)
	// fmt.Println(result1)
	// fmt.Println(result2)

	// demo()
	// fmt.Println(a)
	// println(a)

	// println(math.Pi)
	// var num float64 = 9
	// var result = math.Sqrt(num)
	// fmt.Printf("Result : %.1f", result)

	// IF ELSE

	// num := 20
	// if num == 1 {
	// 	fmt.Println("Hi", num)
	// } else if num == 2 {

	// 	fmt.Println("Bye", num)
	// } else {
	// 	fmt.Println("Okay", num)
	// }

}

func demo() {
	// Local Scope
	var a int = 100
	fmt.Println(a)

}
