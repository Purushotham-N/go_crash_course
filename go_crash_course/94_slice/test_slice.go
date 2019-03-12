package main

import "fmt"

func main() {
	// initialize a slice
	slice := []int{13, 26, 31, 45, 98, 56, 23, 12}

	fmt.Println("slice:", slice)
	fmt.Println("slice[3:8]:", slice[3:8])
	fmt.Println("slice[6:]:", slice[6:])
	fmt.Println("slice[:4]:", slice[:4])
	fmt.Println("slice length:", len(slice))
	fmt.Println("slice length:", cap(slice))

	// modify slice
	for i, v := range slice {
		slice[i] = v - 2
	}
	fmt.Println("The new values in our slice:")
	report(slice)

	// append 2 values
	fmt.Println("Now we will append 2 values slice:")
	slice = append(slice, 7, 8)
	report(slice)

	// add 8 elements to the slice
	slice = resize(slice)
	report(slice)

	//make a copySlice that contains only 8 elemenst
	sliceCopy := make([]int, 8)
	copy(sliceCopy, slice)
	report(sliceCopy)
}

func report(slice []int) {
	fmt.Println("slice:", slice)
	fmt.Println("slice length:", len(slice))
	fmt.Println("slice length:", cap(slice))
}

func resize(slice []int) []int {
	// appending 8 values to the slice
	for i := 0; i < 8; i++ {
		slice = append(slice, 1)
	}
	return slice
}
