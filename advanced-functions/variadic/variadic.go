package variadic

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := addListOfNumbers(&numbers)
	nsum := sumUpN(1, 2, 3, 4, 5, 6)
	dSum := sumUpN(numbers...)
	fmt.Println(sum)
	fmt.Println(nsum)
	fmt.Println(dSum)
}

func addListOfNumbers(numbers *[]int) int {
	sum := 0
	for _, value := range *numbers {
		sum += value
	}
	return sum
}

func sumUpN(numbers ...int) int {
	return addListOfNumbers(&numbers)
}
