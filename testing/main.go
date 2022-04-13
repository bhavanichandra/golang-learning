package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100.0, 0.0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Result of the division is ", result)
}

func divide(x float32, y float32) (float32, error) {
	var result float32 = 0
	if y == 0 {
		return result, errors.New("cannot divide by zero")
	}
	result = x / y
	return result, nil
}
