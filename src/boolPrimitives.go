package main

import (
	"fmt"
)

func main() {
	n := 1 == 1
	fmt.Printf("%v \t % T \n", n, n)

	m := 1 == 2
	fmt.Printf("%v \t %T \n", m, m)

	var f bool = false
	fmt.Printf("%v \t %T \n", f, f)
}
