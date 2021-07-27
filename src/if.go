package main

import (
	"fmt"
)

func literalSyntaxMaps() {
	statePopulations := map[string]int{
		"California":  1234546544,
		"Texas":       25968432,
		"Florida":     577328841,
		"Newyork":     435663235,
		"Pennyslavia": 45473472,
		"Illinois":    3563575365,
		"Ohio":        35646244,
	}
	if pop, ok := statePopulations["Ohio"]; ok {
		fmt.Println(pop)
	}
}

func main() {
	literalSyntaxMaps()
}
