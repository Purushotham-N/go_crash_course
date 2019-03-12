package main

import "fmt"

// this type of functional programming, Recursion
func main() {

	fmt.Println(factorial(5))
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
