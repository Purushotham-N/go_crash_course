package main

import "fmt"

func main() {
	var str [3]string
	str[0] = "Hello"
	str[1] = "My"
	str[2] = "World"

	fmt.Println(str)

	var y [4]int
	y[3] = 10
	fmt.Println(y)

	slice := []int{2, 4, 7, 12, 3}

	fmt.Println(slice)

	

}
