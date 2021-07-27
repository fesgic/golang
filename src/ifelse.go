package main

import (
	"fmt"
	"math"
)

func way1() {
	myNum := 0.01
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are equal")
	} else {
		fmt.Println("These are different")
	}
}

func way2() {
	myNum := 0.01544
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.0001 {
		fmt.Println("These are equal")
	} else {
		fmt.Println("These are different")
	}
}

func main() {
	way1()
	way2()
}
