package main

import (
	"fmt"
	"reflect"
)

type Birds struct {
	Speed   float32
	CanFly  bool
	Habitat string `required:"true" max:"100"`
}

func main() {
	f := reflect.TypeOf(Birds{})         //get type of object we are dealing with
	field, _ := f.FieldByName("Habitat") //get field in object that we want to deal with
	fmt.Println(field.Tag)               //display the tag of the field
}
