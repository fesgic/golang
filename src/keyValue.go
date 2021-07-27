package main

import (
	"fmt"
)

func main() {
	s := []string{"festus", "gichohi", "mungai"}
	for k, v := range s {
		fmt.Println(k, v)
	}
}
