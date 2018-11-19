package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

func makePanic() {

	if user != "liam" {
		panic("no value for user")
	}
}

func throwsPanic(f func()) (b bool) {
	//recover if panic throws
	defer func() {
		//recover return panic msg
		x := recover()
		if x != nil {
			b = true
		}

	}()
	f()
	return
}

func main() {
	fmt.Println(throwsPanic(makePanic))
}
