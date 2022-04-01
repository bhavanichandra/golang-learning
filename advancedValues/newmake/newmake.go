package newmake

import "fmt"

func main() {
	//hobbies := []string{"Sports", "Reading"} //composite literal way of creating slices

	hobbies := make([]string, 2, 10)

	moreHobbies := new([]string)

	fmt.Println(*moreHobbies)

	*moreHobbies = append(*moreHobbies, "Watching")

	fmt.Println(*moreHobbies)

	aMap := make(map[string]int, 5)

	aMap["test"] = 1

	fmt.Println(aMap)

	fmt.Println(hobbies, len(hobbies), cap(hobbies))
	hobbies[0] = "Sports"
	hobbies[1] = "Reading"
	fmt.Println(hobbies)

	hobbies = append(hobbies, "Cooking", "Dancing")

	fmt.Println(hobbies)
}
