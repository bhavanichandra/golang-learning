package exercise

import "fmt"

func ExerciseOne() {
	const PI float64 = 3.14
	radius := 5
	var circumference float64 = 2 * PI * float64(radius)
	var circumference32 float32 = 2 * float32(PI) * float32(radius)

	fmt.Printf("For a radius of %v, the circle circumference (Type: %T) is %.2f \n", radius, circumference, circumference)
	fmt.Printf("For a radius of %v, the circle circumference (Type: %T) is %.2f \n", radius, circumference32, circumference32)

	fmt.Println(PI)
}
