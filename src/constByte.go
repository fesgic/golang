package main

import (
	"fmt"
)

const (
	isAdmin = 1 << iota
	isHeadQuaters
	canSeeFinancials

	overseeAfrica
	overseeAsia
	overseeEurope
	overseeNorthAmerica
	overseeSouthAmerica
)

func main() {
	const admin byte = isAdmin | canSeeFinancials | overseeAfrica | overseeAsia
	fmt.Printf("%b or %v \n", admin, admin)
	fmt.Printf("Can admin oversee Asia?: %v\n", overseeAsia&admin == overseeAsia)
	fmt.Printf("Can admin oversee Europe?: %v", overseeEurope&admin == overseeEurope)
}
