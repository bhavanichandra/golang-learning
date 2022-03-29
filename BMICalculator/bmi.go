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

// calculateBMI calculates BMI for passed weight and height parameters
func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}
