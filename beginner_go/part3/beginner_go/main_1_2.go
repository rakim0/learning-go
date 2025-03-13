package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

// if you want to return multiple objects then
func add_subtract(a, b int) (int, int) {
	return a + b, a - b
}

func main1_2() {
	age := 30
	if age >= 13 {
		fmt.Println("you're a teenager")
	} else if age >= 18 {
		fmt.Println("you're an adult")
	} else {
		fmt.Println("you are a child")
	}

	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week!")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Midweek!")
	case "Friday":
		fmt.Println("TGIF")
	default:
		fmt.Println("Weekend!")
	}

	// for loop
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// while loop
	counter := 0
	for counter < 3 {
		fmt.Println("this is the counter", counter)
		counter++
	}

	// infinite loop
	// for {

	// }

	// arrays
	numbers := [5]int{1, 2, 3, 4, 5}
	for i := range len(numbers) {
		fmt.Println(numbers[i])
	}

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	// allNumbers := numbers[:]
	// firstThree := numbers[0:3]

	fruits := []string{}
	fruits = append(fruits, "kiwi")
	fmt.Println("Fruits: ", fruits)
	fruits = append(fruits, "blueberry", "orange")
	fmt.Println("Fruits: ", fruits)
	moreFruits := []string{"apple", "pineapple"}
	fruits = append(fruits, moreFruits...)
	fmt.Println("Fruits: ", fruits)

	// making a slice with an initial size
	numbers2 := make([]int, 2)                 // slice with length 2 and capacity 2
	numbersWithCapacity := make([]int, 2, 100) // slice with length 2 and capacity 3
	fmt.Println(len(numbersWithCapacity), cap(numbersWithCapacity))
	fmt.Println(len(numbers2), cap(numbers2))
	fmt.Println("Uninitialized slice: ", numbersWithCapacity[0])
	// iterating over values
	fmt.Println("Numbers array: ")
	for _, val := range numbers {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	capitalCities := map[string]string{
		"US":    "Washington D.C.",
		"India": "New Delhi",
		"UK":    "London",
	}

	capital, exists := capitalCities["Germany"]
	if exists {
		fmt.Printf("Capital: %s", capital)
	} else {
		fmt.Printf("Capital of Germany doesn't exist in the map.\n")
	}
	delete(capitalCities, "UK")
	fmt.Printf("%v\n", capitalCities)
}
