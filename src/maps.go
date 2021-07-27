package main

import (
	"fmt"
)

func literalSyntax(){
	statePopulations := map[string]int{
		"California" : 1234546544,
		"Texas" : 25968432,
		"Florida" : 577328841,
		"Newyork" : 435663235,
		"Pennyslavia" : 45473472,
		"Illinois" : 3563575365,
		"Ohio" : 35646244,
	}
	fmt.Println(statePopulations)
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations);
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
}

func makeSyntax(){
	//used when declaring variables and you
	//dont have entries that you are gonna put to it
	countryPopulations := make(map[string]int)
	countryPopulations = map[string]int{
		"Kenya" : 46324353,
		"Uganda" : 3467344,
		"Tanzania" : 56774545,
		"Somalia" : 4546534,
	}
	fmt.Println(countryPopulations)
}
func main(){
	literalSyntax()
	var m = map[[1]string]int{
		["festus"]: 2
	}
	fmt.Println("\n",m,"\n")
	makeSyntax()

}