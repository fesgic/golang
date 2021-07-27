package main

import (
	"fmt"
)

func main() {
	a := 10 //1010
	b := 3  //0011
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(float32(a) / float32(b))
	fmt.Println(a%b, "\n")

	fmt.Println(a & b)
	fmt.Println(a / b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
}
