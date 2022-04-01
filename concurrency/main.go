package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().Unix())
var randN = rand.New(source)

func main() {

	var x = make(chan int)
	var y = make(chan int)

	limiter := make(chan int, 3)

	go generateValue(x, limiter)
	go generateValue(y, limiter)

	var a int
	var b int

	select {
	case a = <-x:
		fmt.Printf("X finished. Value from x is: %v\n", a)
	case b = <-y:
		fmt.Printf("Y finished. Value from y is: %v\n", b)
	}

	// go generateValue(ch, limiter)
	// go generateValue(ch, limiter)

	// sum := 0
	// i := 0

	// for num := range ch {
	// 	sum += num
	// 	i++
	// 	if i == 4 {
	// 		close(ch)
	// 	}
	// }

}

func generateValue(c chan int, limit chan int) int {
	limit <- 1
	fmt.Println("Generating Value")
	sleepTime := randN.Intn(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)
	random := randN.Intn(10)
	c <- random
	<-limit
	return random
	// close(c) => closes the channel
}
