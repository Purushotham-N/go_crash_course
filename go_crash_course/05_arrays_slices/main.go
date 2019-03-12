package main

import "fmt"

func main() {
	// Arrays
	// var fruitArray [2]string

	// Assign values
	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"

	// declare and assign
	// fruitArray := [2]string{"Apple", "Orange"}

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:])
}
