package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	Speed  float32
	CanFly bool
}

func positionalSyntaxComposition() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.Speed = 48
	b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Name, "\n"+b.Origin+"\n", b.Speed, "\n", b.CanFly)
}

func literalSyntaxEmbedding() {
	c := Bird{Animal: Animal{Name: "Chicken", Origin: "All Continents"},
		Speed:  20,
		CanFly: false,
	}
	fmt.Println(c)
	fmt.Println(c.Name, "\n", c.Origin, "\n", c.Speed, "\n", c.CanFly)
}

func main() {
	positionalSyntaxComposition()
	fmt.Print("\n")
	literalSyntaxEmbedding()
}
