package main

import "fmt"

func main() {
	defer fmt.Println("Counting..")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Finished..")
}
