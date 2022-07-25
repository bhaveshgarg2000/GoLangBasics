package main

import "fmt"

func main() {

	arr := []int{10, 20, 30, 40, 50}
	count := 0

	for i := 0; i < len(arr); i++ {
		count += arr[i]
	}
	fmt.Println(count)
}
