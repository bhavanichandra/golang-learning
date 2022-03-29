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

func getUserMetrics() (float64, float64) {

	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')
	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	weight := parseFloat(weightInput)
	height := parseFloat(heightInput)

	return weight, height
}

func parseFloat(input string) float64 {
	convertedInput := strings.Replace(input, "\r\n", "", -1)
	parsedInput, _ := strconv.ParseFloat(convertedInput, 64)
	return parsedInput
}
