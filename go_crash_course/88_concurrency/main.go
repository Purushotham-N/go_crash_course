package main

import "fmt"

func printMsg(msg string) {
	for i := 0; i < 20; i++ {
		fmt.Println(msg, ":", i)
	}
}

func main() {
	go printMsg("value of i")
	var input string
	fmt.Println("Hello world..!")
	fmt.Scanf("enter a string:" , &input)
}
