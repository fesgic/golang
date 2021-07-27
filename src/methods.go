package main

import (
	"fmt"
)

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println("The name is:", g.name)
}

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	//pass by pointers to prevent duplicating when working with large datasets and
	//change value in struct in the function exception with slices and maps since
	//they refer to original data and are not copy types but memory types
	fmt.Println(g.greeting, g.name) //methods in go
	g.name = "Festus"
}
