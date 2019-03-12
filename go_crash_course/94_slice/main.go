package main

import "fmt"

func main() {
	slice := []int{1, 2, 4, 8, 16, 32, 64, 128}
	fmt.Println("Slice is: ", slice)

	for i := 0; i < len(slice); i++ {
		fmt.Printf("silce[%d] == %d  \n", i, slice[i])
	}

	// slicing
	fmt.Printf("slice[2:4] == %d\n", slice[2:4])
	fmt.Printf("slice[:5] == %d\n", slice[:5])
	fmt.Printf("slice[3:] == %d\n", slice[3:])
	fmt.Printf("slice[:] == %d\n", slice[:])

	// calling getPow
	fmt.Println("Slice is: ", slice)
	getPower(slice)
	makeSlice()
}

func getPower(slice []int) {
	for n, p := range slice {
		fmt.Printf("2**%d == %d\n", n, p)
	}
}

func makeSlice() {

	a := make([]int, 4)
	printSl("a", a)
	b := make([]int, 0, 4)
	printSl("b", b)
	c := b[:1]
	printSl("c", c)
	d := c[2:4]
	printSl("d", d)
}

func printSl(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
