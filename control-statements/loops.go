package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	userChoice, err := getUserChoice()
	if err != nil {
		fmt.Println("Invalid Choice. Exiting..!")
		return
	}

	if userChoice == "1" {
		calculateSumOfNNumbers()
	} else if userChoice == "2" {
		calculateFactorial()
	} else if userChoice == "3" {
		calculateManuallyEnteredSum()
	} else if userChoice == "q" {
		fmt.Println("Quitting...!")
		return
	} else {
		calculateListSum()
	}
}

func calculateSumOfNNumbers() {
	fmt.Print("Please enter a number: ")
	chosenNumber, err := getInputNumber()
	if err != nil {
		fmt.Println("Invalid number input")
		return
	}
	sum := 0
	for i := 1; i <= chosenNumber; i++ {
		sum = sum + i
	}

	fmt.Printf("Result: %v\n", sum)

}

func calculateFactorial() {
	fmt.Print("Please enter a number: ")
	chosenNumber, err := getInputNumber()
	if err != nil {
		fmt.Println("Invalid number input")
		return
	}
	product := 1
	for i := 1; i <= chosenNumber; i++ {
		product = product * i
	}

	fmt.Printf("Result: %v\n", product)
}

func calculateManuallyEnteredSum() {
	isEnteringNumbers := true
	sum := 0
	fmt.Print("Keep on entering numbers, the program will exit once any other input is entered ")
	for isEnteringNumbers {
		inputNumber, err := getInputNumber()
		sum = sum + inputNumber
		isEnteringNumbers = err == nil
	}
	fmt.Printf("result: %v\n", sum)
}

func calculateListSum() {
	fmt.Println("Please enter a comma-separated list of numbers")
	inputNumberList, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid Input")
		return
	}
	inputNumberList = cleanInput(inputNumberList)
	inputNumbers := strings.Split(inputNumberList, ",")
	sum := 0
	for index, value := range inputNumbers {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
		number, err := strconv.ParseInt(value, 0, 64)
		if err != nil {
			fmt.Printf("Skipping value: %v | Reason: %v", value, err)
			continue
		}
		sum = sum + int(number)
	}
	fmt.Printf("result: %v\n", sum)
}

func getUserChoice() (string, error) {
	fmt.Println("Operations")
	fmt.Println("---------------------------------")
	fmt.Println("1) Add all numbers upto number x")
	fmt.Println("2) Build factorial upto number x")
	fmt.Println("3) sum up manually entered numbers")
	fmt.Println("4) sum up list of entered numbers")
	fmt.Println("q) Quit the program")

	fmt.Print("Please make your choice: ")
	userChoice, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	userChoice = cleanInput(userChoice)
	if userChoice == "1" || userChoice == "2" || userChoice == "3" || userChoice == "4" || userChoice == "q" {
		return userChoice, nil
	} else {
		err = errors.New("INVALID_OPTION")
		return "", err
	}

}

func cleanInput(input string) string {
	if runtime.GOOS == "windows" {
		return strings.Replace(input, "\r\n", "", -1)
	} else {
		return strings.Replace(input, "\n", "", -1)
	}
}

func getInputNumber() (int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = cleanInput(input)
	number, err := strconv.ParseInt(input, 0, 64)
	if err != nil {
		return 0, err
	}
	return int(number), nil
}
