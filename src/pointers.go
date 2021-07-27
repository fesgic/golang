package main

import "fmt"

func pointArray() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v, %p, %p", a, b, c)
}

type myStruct struct {
	foo int
}

func pointerStruct() {
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)
}

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, *b, b, &b)
	a = 24
	fmt.Println(a, *b, b, &b, "\n")
	pointArray()
	fmt.Println("\n")
	pointerStruct()
}
