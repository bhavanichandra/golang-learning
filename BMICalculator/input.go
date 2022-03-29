package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bhavanichandra/bmi/info"
)

var reader = bufio.NewReader(os.Stdin)

// getUserMetrics function returns user entered metrics
func getUserMetrics() (float64, float64) {

	fmt.Print(info.WeightPrompt)
	weight := getInput()
	fmt.Print(info.HeightPrompt)
	height := getInput()

	return weight, height
}

// getInput finction, fetches input from user, returns the parsed data
func getInput() float64 {
	input, _ := reader.ReadString('\n')
	convertedInput := strings.Replace(input, "\r\n", "", -1)
	parsedInput, _ := strconv.ParseFloat(convertedInput, 64)
	return parsedInput
}
