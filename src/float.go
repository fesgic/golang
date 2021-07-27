package main

import (
	"fmt"
)

func main() {
	var n float64 = 3.142
	n = 13.7e72
	n = 2.1e14
	fmt.Printf("%g %[1]T", n)
}
