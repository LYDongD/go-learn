package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)
	o := make(chan bool)

	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout > 5s")
				o <- true
				break
			}
		}
	}()

	//blocked until timeout, o being writtern true
	<-o
}
