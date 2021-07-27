package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	fmt.Print(fib(8), "\n")
	f, err := ioutil.ReadFile("src/foo.txt")
	if err != nil {
		log.Fatal(err)
	}
	file := string(f)
	fmt.Print(file)

}
