package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, getTransformerFunction(&numbers))
	tripled := transformNumbers(&doubled, getTransformerFunction(&doubled))

	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	var dNumbers []int
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}
