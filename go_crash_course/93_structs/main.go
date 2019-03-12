package main

import "fmt"

type Rectangle struct {
	width, height int
}

func main() {
	fmt.Println(Rectangle{7, 8})

	r := Rectangle{1, 2}
	r.width = 18
	fmt.Println(r.width)

	s := Rectangle{4, 6}

	p := &s

	p.width = 8
	fmt.Println(s)

	r1 := Rectangle{7, 8}       // type Rectangle with width=7, height=8
	r2 := Rectangle{width: 22}  // type Rectangle with width=22, height=0 implicitely
	r3 := Rectangle{height: 44} // type Rectangle with width=0 implicitly, height=44
	r4 := Rectangle{}           // type Rectangle with width=0, height=0 implicitly
	z := &Rectangle{7, 8}       // type * Rectangle

	fmt.Println(r1, r2, r3, r4, z)
}
