package main

import (
	"fmt"
)

//type omit
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//multi return
func sumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

//multi params
func multiParamsFunc(args ...int) {
	for _, num := range args {
		fmt.Print(num)
	}
}

//pass point
func add(a *int) int {
	*a = *a + 1
	return *a
}

//func main() {
//
//	//fmt.Println(max(1, 3))
//
//	//x, y := sumAndProduct(1, 3)
//	//fmt.Println(x, y)
//
//	//multiParamsFunc(1, 2, 3, 4, 5)
//
//	a := 1
//	add(&a)
//	fmt.Println(a)
//}
