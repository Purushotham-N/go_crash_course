package main

import "fmt"

type Rect struct {
	width, height int
}

var m = map[string]Rect{
	"Rect1": Rect{1, 2},
	"Rect2": Rect{4, 6},
}

func main() {
	fmt.Println(m)

	w := make(map[string]int)
	w["Answer"] = 10
	fmt.Println("The value:", w["Answer"])

	w["Answer"] = 20
	fmt.Println("The value:", w["Answer"])

	delete(w, "Answer")
	fmt.Println("The value:", w["Answer"])

	v, ok := w["Answer"]
	fmt.Println("The value:", v, "present?", ok)

}
