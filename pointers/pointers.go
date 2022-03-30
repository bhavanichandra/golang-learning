package main

import "fmt"

func main() {
	age := 26
	fmt.Println(age)

	//var myAge *int
	//myAge = &age
	myAge := &age

	fmt.Printf("age %v stored at %v.\n", age, myAge)

	*myAge = 33 // Using * on a pointer variable -> its called dereference
	fmt.Printf("age %v stored at %v.\n", age, myAge)

	doubledAge := double(myAge)
	fmt.Println(doubledAge)
	fmt.Println(age)
}

func double(number *int) int {
	result := *number * 2
	*number = 25
	return result
}
