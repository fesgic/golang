package main

import (
	"fmt"
)

func main() {
	s := "this is a string"
	b := []byte(s)
	fmt.Printf("%v %T", b, b)
}
