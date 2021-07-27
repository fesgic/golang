package main

import (
	"fmt"
)

var a float32 = 21

func main() {
	var i int
	i = 42
	var j int = 56
	k := 27
	fmt.Println(i, " ", j, " ", k)
	fmt.Printf("%v %T", j, j)
	fmt.Printf("    %v %T", a, a)
}
