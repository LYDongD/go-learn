package main

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box

func (b Box) Volumn() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, box := range bl {
		if bv := box.Volumn(); bv > v {
			v = bv
			k = box.color
		}
	}

	return k
}

func (bl BoxList) PaintBlack() {
	for i := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {

	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
	}

	fmt.Println(boxes[0].Volumn())
	fmt.Println(boxes[len(boxes)-1].color.String())
	fmt.Println(boxes.BiggestColor())

	boxes.PaintBlack()
	fmt.Println(boxes[0].color.String())
	fmt.Println(boxes.BiggestColor())
}
