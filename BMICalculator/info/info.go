package info

import "fmt"

const mainTitle = "BMI Calculator"
const separator = "--------------------------"
const WeightPrompt = "Please enter your weight (kg): "
const HeightPrompt = "Please enter your height (m): "

// PrintWelcome function prints welcome to user
func PrintWelcome() {
	fmt.Println(mainTitle)
	fmt.Println(separator)
}
