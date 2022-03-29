package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a, b := generateRandomNumbers()
	sum := add(a, b)
	fmt.Printf("%v + %v = %v \n", a, b, sum)
	printNum(sum)
}

func generateRandomNumbers() (r1 int, r2 int) {
	source := rand.NewSource(time.Now().Unix())
	rnd1 := rand.New(source)
	r1 = rnd1.Intn(148)
	r2 = rnd1.Intn(289)
	return
}

// add function adds two integers and returns an integer
func add(num1 int, num2 int) (sum int) {
	sum = num1 + num2
	return
}

func printNum(value int) {
	fmt.Println(value)
}
