package main

import (
	"fmt"
	"github.com/bhavanichandra/firstproject/exercise"
	"github.com/bhavanichandra/firstproject/greeting"
)

func main() {
	//var greetingText string   // Long way to declare
	//greetingText = "Hello World" // Assign value

	// var greetingText string = "Hello World" // Another way to declare and assign value to variable

	//greetingText2 := "Hello World" // Shortest way to declare and assign variable
	luckyNumber := 17 // int support any number with limitations at very large number. Only Integers, not decimals

	eventMoreLuckyNumber := luckyNumber + 5

	fmt.Println(greeting.GreetingText)
	fmt.Println(greeting.GreetingText)
	fmt.Println(luckyNumber)
	fmt.Println(eventMoreLuckyNumber)

	var newNumber float64 = float64(luckyNumber) / 3

	fmt.Println(newNumber)

	var defaultFloat float64 = 1.2367823678923678236789
	var smallFloat float32 = 1.2367823678923678236789

	fmt.Println(defaultFloat)
	fmt.Println(smallFloat)

	var firstRune rune = '$'
	fmt.Println(firstRune)
	fmt.Println(string(firstRune))

	var firstByte byte = 's'
	fmt.Println(firstByte)
	fmt.Println(string(firstByte))

	firstName := "Bhavani Chandra"
	lastName := "Vajapeyayajula"
	//fullName := fmt.Sprintln(firstName, " ", lastName)
	fullName := fmt.Sprintf("%v %v", firstName, lastName)
	age := 26

	fmt.Printf("Hi, I'm %v and I'm %v (Type: %T) years old", fullName, age, age)
	//fmt.Println("9" + 1) // Mixed type

	fmt.Println("\n-------Exercise One-----------")
	exercise.ExerciseOne()
	fmt.Println("---------Exercise Two----------")
	exercise.ExerciseTwo()
	fmt.Println("-------------------------------")
}
