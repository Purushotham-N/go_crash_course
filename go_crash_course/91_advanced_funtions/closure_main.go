package main

import "fmt"

func main() {
	a := 4
	// this type of defining inner funtions is called as closure
	decrement := func() int {
		a--
		return a
	}
	fmt.Println(decrement())
	fmt.Println(decrement())
}
