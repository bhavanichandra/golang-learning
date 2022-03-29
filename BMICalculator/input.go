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
func getUserMetrics() (weight float64, height float64) {
	weight = getInput(info.WeightPrompt)
	height = getInput(info.HeightPrompt)
	return
}

// getInput function, fetches input from user, returns the parsed data
func getInput(promptText string) float64 {
	fmt.Print(promptText)
	input, _ := reader.ReadString('\n')
	convertedInput := strings.Replace(input, "\r\n", "", -1)
	parsedInput, _ := strconv.ParseFloat(convertedInput, 64)
	return parsedInput
}
