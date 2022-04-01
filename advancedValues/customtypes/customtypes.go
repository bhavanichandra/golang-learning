package customtypes

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type customNumber int

func (number customNumber) pow(power int) customNumber {
	result := number
	for i := 1; i < power; i++ {
		result *= number
	}
	return result
}

type personData map[string]person

func main() {
	var persons personData = personData{
		"p1": {"bhavani", 27},
	}
	fmt.Println(persons)

	var dummyNumber customNumber = 5
	power := dummyNumber.pow(2)
	fmt.Println(power)
}
