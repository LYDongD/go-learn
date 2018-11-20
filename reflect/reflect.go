package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.27
	v := reflect.ValueOf(x)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind: ", v.Kind())
	fmt.Println("value: ", v.Float())

	p := reflect.ValueOf(&x)
	v = p.Elem()
	v.SetFloat(7.1)
	fmt.Println("value: ", v.Float())

}
