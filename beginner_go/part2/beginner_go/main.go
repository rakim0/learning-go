package main

import (
	"fmt"
)

func main() {
	var name string = "Rakim"
	fmt.Printf("This is my name: %s\n", name)

	age := 21
	fmt.Printf("My age is %d\n", age)

	var city string
	fmt.Printf("My city is %s\n", city)

	city = "Prayagraj"

	fmt.Printf("My city is %s\n", city)

	var country, continent string = "India", "Asia"
	fmt.Printf("My country is %s and my continent is %s\n", country, continent)

	var (
		isEmployeed bool   = false
		salary      int    = 40000
		position    string = "developer"
	)
	fmt.Printf("isEmployeed: %t, Salary: %d, Position: %s\n", isEmployeed, salary, position)
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool
	fmt.Printf("int: %d float: %f string: %s bool: %t\n", defaultInt, defaultFloat, defaultString, defaultBool)

	const (
		jan = iota + 1
		feb
		mar
		apr
	)
	fmt.Printf("Jan - %d Feb - %d Mar - %d Apr - %d\n", jan, feb, mar, apr)
	result := add(1, 2)
	fmt.Println("this is the result", result)
	sum, diff := add_subtract(1, 2)
	fmt.Printf("Sum: %d and Difference: %d\n", sum, diff)
}

func add(a int, b int) int {
	return a + b
}

// if you want to return multiple objects then
func add_subtract(a, b int) (int, int) {
	return a + b, a - b
}
