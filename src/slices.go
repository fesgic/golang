package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(a, "\n")
	b := a[1:]
	fmt.Println(b, "\n")
	c := a[:len(a)-1]
	fmt.Println(c)
	d := append(a[:3], a[4:]...)
	fmt.Println(d)
	fmt.Println(a)
}
