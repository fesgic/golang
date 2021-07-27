package main

import (
	"fmt"
)

func switchCase1() {
	switch 1 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Print("Two")
	default:
		fmt.Println("Not one or two")
	}
}

func switchCase2() {
	switch i := 2 + 3; i {
	case 1, 2, 3, 4, 5:
		fmt.Println("One,two,three, four,five")
	case 6, 7, 8, 9, 10:
		fmt.Println("Six,seven,eight,nine,ten")
	default:
		fmt.Println("Number not included")
	}
}

func switchCase3() {
	i := 10
	switch {
	case i <= 10:
		fmt.Println("Less than or equal to ten")
	case i <= 20:
		fmt.Println("Less than or equal to twenty")
	default:
		fmt.Println("Greater than ten")
	}
}

func typeCase() {
	var i interface{} = []string{}
	switch i.(type) {
	case int:
		fmt.Println("This is an integer")
	case []string:
		fmt.Println("This is a slice")
	case float64:
		fmt.Println("This is a float64")
	case bool:
		fmt.Println("This is a boolean")
	case complex128:
		fmt.Println("This is a complex64")

	case string:
		fmt.Println("This is a string")
	default:
		fmt.Println("Type is not defined")
	}
}

func main() {
	switchCase1()
	switchCase2()
	switchCase3()
	typeCase()
}
