package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	//Using var
	//var name string = "Purushotham"
	var age int32 = 32
	const isCool = true
	var size float32 = 2.5

	//shorthand
	// name := "Purushotham"
	// email := "test@gmail.com"
	// size := 1.35

	name, email := "brad", "brad@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T", size)
}
