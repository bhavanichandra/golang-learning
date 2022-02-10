package exercise

import "fmt"

func ExerciseTwo() {
	var firstName = "Bhavani Chandra"
	lastName := "Vajapeyayajula"

	fmt.Println(firstName)
	fmt.Println(lastName)

	fullName := firstName + " " + lastName

	fmt.Println(fullName)

	var currentYear int = 2022
	birthYear := 1995

	ageDifference := currentYear - birthYear

	fmt.Println(ageDifference)

	currentYear = currentYear + 1
	ageDifference = currentYear - birthYear

	fmt.Println(ageDifference)
}
