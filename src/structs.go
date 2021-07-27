package main

import (
	"fmt"
)

type doctor struct {
	number     int
	actorNames string
	companions []string
}

func literalSyntaxStruct() {
	aDoctor := doctor{
		number:     1,
		actorNames: "John Smith",
		companions: []string{
			"John Pertwee",
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.number, aDoctor.actorNames)
	for x := 0; x < len(aDoctor.companions); x++ {
		fmt.Println(x+2, aDoctor.companions[x])
	}
}

func positionalSyntaxStruct() {
	bDoctor := doctor{
		2,
		"Festus Gichohi",
		[]string{
			"Vincent Nderitu",
			"Roseanne Maina",
			"Faith Kate",
			"Faith Njoroge",
			"Joseph Kamau",
		},
	}
	fmt.Println(bDoctor)
	fmt.Println(bDoctor.number, bDoctor.actorNames)
	for x := 0; x < len(bDoctor.companions); x++ {
		fmt.Println(bDoctor.companions[x])
	}
}

func anonymousSyntaxStruct() {
	cDoctor := struct {
		name       string
		companions []string
	}{"John Mwangi", []string{"Bob Njoroge", "Festus Gichohi"}}
	fmt.Println(cDoctor)
	for x := 0; x < len(cDoctor.companions); x++ {
		fmt.Println(cDoctor.companions[x])
	}
}

func main() {
	literalSyntaxStruct()
	fmt.Println("\n")
	positionalSyntaxStruct()
	fmt.Println("\n")
	anonymousSyntaxStruct()
}
