package main

import "fmt"

// this type of functional programming, fabonnocci closure
func fibonacci() func() int {
	x := 0
	y := 1
	fmt.Println(x)
	return func() int {
		x, y = y, x+y
		return x
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
