package main

import (
	"fmt"
	"math"
)

func hypot(x, y float64) float64 {
	return math.Pow(y*y+x*x, 2)
}

func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

func divide(b, d float64) (float64, error) {
	if d == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return b / d, nil
}

func main() {
	fmt.Println(hypot(2, 3))

	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is:", *s)

	e, err := divide(2.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(e)

	//fmt.Println(divide(2,0))
}
