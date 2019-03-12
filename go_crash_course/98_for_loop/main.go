package main

import "fmt"

// func main() {
// 	z := 1
// 	for z < 15 {
// 		fmt.Println(z)
// 		z = z + 1
// 	}
// }

func main() {
	for z := 1; z <= 10; z++ {
		defer fmt.Println(z)
	}
	fmt.Println("Finished..")
}
