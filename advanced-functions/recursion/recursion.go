package recursion

import "fmt"

func main() {
	num := 5
	result := factorial(num)
	fmt.Printf("%v! = %v\n", num, result)
}

func factorial(num int) int {
	if num <= 1 {
		return 1
	}

	return num * factorial(num-1)
}
