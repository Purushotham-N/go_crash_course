// Go use a very simple way to solve it. The outer fields get upper access levels, which means when you access student.phone,
// we will get the field called phone in student, not the one in the Human struct.
// This feature can be simply seen as field overloading.

package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string // Human has phone field
}

type Employee struct {
	Human
	specialty string
	phone     string // phone in employee
}

func main() {
	Bob := Employee{Human{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}

	fmt.Println("Bob's work phone is:", Bob.phone)
	fmt.Println("Bob's personal phone is:", Bob.Human.phone)
}
