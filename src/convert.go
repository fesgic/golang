package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	fmt.Printf("%v \t %T \n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v \t %T \n", j, j)

	k := 4
	fmt.Printf("%v \t %T \n", k, k)
}
