package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	Skills
	int
	weight     int
	speciality string
}

func (h *Human) SayHi() {
	fmt.Printf("i am %s, %d this year\n", h.name, h.age)
}

func main() {

	//inherit
	mark := Student{Human: Human{"mark", 22, 100}, weight: 120, speciality: "math"}
	fmt.Println(mark.name)
	fmt.Println(mark.age)
	//override
	fmt.Println(mark.Human.weight)
	fmt.Println(mark.weight)
	fmt.Println(mark.speciality)

	mark.Skills = []string{"music", "basketball"}
	fmt.Println(mark.Skills)
	mark.Skills = append(mark.Skills, "speech", "golang")
	fmt.Println(mark.Skills)

	mark.int = 3
	fmt.Println(mark.int)

	mark.SayHi()
}
