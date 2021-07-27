package main

import (
	"fmt"
)

func main() {
	a := "r"
	c := []byte(a)
	b := 'r'
	fmt.Printf("%v %T \n", c, c)
	fmt.Printf("%v %T", b, b)
}
