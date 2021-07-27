package main

import (
	"fmt"
	"math"
)

func simpleLoop() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d , e = %8.3f\n", x, math.Exp(float64(x)))
	}
}

func doubleLoop() {
	for i, j := 1, 2; j <= 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
}

func main() {
	fmt.Println("Simple loop")
	simpleLoop()
	fmt.Println("\nDouble loop")
	doubleLoop()
}
