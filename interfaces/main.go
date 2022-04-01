package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

type logger interface {
	log()
}

type logData struct {
	message  string
	fileName string
}

func (lData *logData) log() {
	err := ioutil.WriteFile(lData.fileName, []byte(lData.message), fs.FileMode(0644))

	if err != nil {
		fmt.Println("Storing the log data failed!")
	}
}

type loggableString string

func (text loggableString) log() {
	fmt.Println(text)
}

func main() {
	userLog := &logData{"User logged in", "user-log.txt"}
	// do more work ...
	userLog.log()
	//createLog(userLog) => Not work as createLog takes loggableString type as parameter
	createLog(userLog)

	message := loggableString("[INFO] User created!")
	// do more work ...
	createLog(message)
	outputValue(message)
	outputValue(userLog)

	firstNumber := 5
	secondNumber := 23.9
	numbers := []int{1, 2, 3}

	doubledFirstNumber := double(firstNumber)
	doubledSecondNumber := double(secondNumber)
	doubledNumbers := double(numbers)
	doubledString := double("Test")
	//createLog("Random String") -> This cannot work of createLog takes logger

	fmt.Println(doubledFirstNumber)
	fmt.Println(doubledSecondNumber)
	fmt.Println(doubledNumbers)
	fmt.Println(doubledString)
}

func createLog(data logger) {
	// More things to do...
	data.log()
}

func outputValue(value interface{}) { // Empty interface, which enables to have any value
	//val, ok := value.(loggableString) => Type Assertion
	//fmt.Println(val, ok)
	fmt.Println(value)
}

func double(value interface{}) interface{} {
	switch val := value.(type) { // Type Switch
	case int:
		return val * 2
	case float64:
		return val * 2
	case []int:
		newNumber := append(val, val...)
		return newNumber
	default:
		return ""
	}
}
