package main

import (
	"github.com/bhavanichandra/bmi/info"
)

func main() {
	// Output information
	info.PrintWelcome()
	weight, height := getUserMetrics()
	//Calculate BMI
	bmi := calculateBMI(weight, height)
	// Output BMI
	printBMI(bmi)
}

func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}
