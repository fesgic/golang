package main

import (
	"fmt"
)

func main() {
	p := new(int) //p, of type int, points value to unnamed T variable
	fmt.Print(*p) //is usually defaulted to zero
	*p = 1
	fmt.Print("\n", *p, "\n")

	fmt.Print(*newInt(), "\n")
	fmt.Print(*oldInt())
}

func newInt() *int {
	return new(int)
}

func oldInt() *int {
	var dummy int
	return &dummy
}
