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
	weight := getInput()
	fmt.Print(info.HeightPrompt)
	height := getInput()

	return weight, height
}

func getInput() float64 {
	input, _ := reader.ReadString('\n')
	convertedInput := strings.Replace(input, "\r\n", "", -1)
	parsedInput, _ := strconv.ParseFloat(convertedInput, 64)
	return parsedInput
}
