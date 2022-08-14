// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var i = 15
// 	var txt = "Hello World!"

// 	fmt.Printf("%v\n", i)
// 	fmt.Printf("%#v\n", i)
// 	fmt.Printf("%v%%\n", i)
// 	fmt.Printf("%T\n", i)

// 	fmt.Printf("%v\n", txt)

// 	fmt.Printf("%#v\n", txt)
// 	fmt.Printf("%T\n", txt)
// }

package main

import (
	"fmt"
)

func main() {
	var txt = "Hello"
	var i = 3.14

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)
	fmt.Printf("%g\n", i)
}
