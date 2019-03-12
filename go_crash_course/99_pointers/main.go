package main

import "fmt"

func fun(y int) {
	y = 0
}

func funPoint(yPnt *int) {
	*yPnt = 0
}

func intiateValue(pnt *int) {
	*pnt = 99
}
func main() {
	y := 10
	fmt.Println("Before passing to fun:y=", y)
	fun(y)
	fmt.Println("After passing to fun:y=", y)
	funPoint(&y)
	fmt.Println("After passing to funPoint:y=", y)

	// create a addess location with size
	zPnt := new(int)
	intiateValue(zPnt)
	fmt.Println("After passing to intiateValue:zPnt=", zPnt)
	fmt.Println("After passing to intiateValue:*zPnt=", *zPnt)
}
