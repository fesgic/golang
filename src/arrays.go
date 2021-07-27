package main

import (
	"fmt"
)

func array1() {
	var students [2]string
	students[0] = "Festus"
	students[1] = "Gichohi"
	fmt.Printf("Students: %v\n", students)
	for x := 0; x < len(students); x++ {
		fmt.Printf("Student #%v \t %v\n", x, students[x])
	}
}
func main() {
	grades := [3]int{97, 45, 23}
	fmt.Printf("Grades: %v\n", grades)
	array1()
	array2()
}

func array2() {
	a := [...]int{1, 2, 3}
	b := &a
	b[1] = 25
	fmt.Println(a)
	fmt.Println(*b)
}
