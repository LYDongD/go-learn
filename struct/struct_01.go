package main

import "fmt"

type person struct {
	name string
	age  int
}

func Older(p1, p2 person) (person, int) {

	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}

	return p2, p2.age - p1.age
}

func main() {

	var Liam person
	Liam.name = "liam"
	Liam.age = 28

	Pony := person{name: "pony", age: 29}

	miaomei := new(person)
	miaomei.name = "miao"
	miaomei.age = 4

	p, diff := Older(Liam, Pony)
	fmt.Printf("older: %s, differ:%d\n", p.name, diff)
	fmt.Println(miaomei.name, miaomei.age)
}
