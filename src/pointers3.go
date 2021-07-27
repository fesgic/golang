package main

import "fmt"


func inc(*int) int{
	*p++
	return p
}
var p = 1;
inc(&p)

func main(){
	inc(&p);
	fmt.Println(inc(&p))
}