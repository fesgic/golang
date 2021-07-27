package main

import (
	"fmt"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	filesize := 4857494878
	fmt.Println("Size is", filesize/GB, "GB")
	fmt.Printf("%v GB", float32(filesize)/float32(GB))

}
