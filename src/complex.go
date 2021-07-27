package main

import (
	"fmt"
)

func realimag() {
	var a complex64 = 1 + 7i
	fmt.Printf("%v %T\n", real(a), real(a))
	fmt.Printf("%v %T\n \n", imag(a), imag(a))

	var b complex128 = 1 + 14i
	fmt.Printf("%v %T\n", real(b), real(b))
	fmt.Printf("%v %T\n\n", imag(b), imag(b))
}

func complexno() {
	var c complex128 = complex(5, 12)
	fmt.Printf("%v %T", c, c)
}
func main() {
	n := 1 + 2i
	fmt.Printf("%v %T \n \n", n, n)
	m := 2 + 5.2i
	fmt.Println(n + m)
	fmt.Println(n - m)
	fmt.Println(n * m)
	fmt.Println(n/m, "\n \n")

	realimag()
	complexno()

}
