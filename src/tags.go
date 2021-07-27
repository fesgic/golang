package main

import (
	"fmt"
	"reflect"
)

type Animals struct {
	Name   string `required max:"100"`
	Origin string
}

func main() {
	t := reflect.TypeOf(Animals{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
