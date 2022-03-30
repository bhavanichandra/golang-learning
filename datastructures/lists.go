package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func NewProduct(id string, title string, price float64) *Product {
	return &Product{id, title, price}
}

func main() {
	hobbies := [3]string{"Reading", "Swimming", "Cycling"}

	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	//mainHobbies := hobbies[0:2]
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	goals := []string{"Learn", "Enjoy"}
	goals[1] = "Be Happy"
	goals = append(goals, "Be Kind")
	fmt.Println(goals)

	products := []Product{
		*NewProduct("a-book", "A great book", 25.99),
		*NewProduct("a-carpet", "Great Carpet", 29.99),
	}

	products = append(products, *NewProduct("another-book", "Another great book", 59.99))
	fmt.Println(products)

}

func courseContent() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices, prices)
}

//func main() {
//	var productNames = [4]string{"A Book"}
//	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
//	fmt.Println(prices)
//	productNames[2] = "A Carpet"
//	fmt.Println(productNames)
//	fmt.Println(prices[2])
//
//	featuredPrices := prices[1:]
//	featuredPrices[0] = 12.99
//	highlightedPrices := featuredPrices[:1]
//	fmt.Println(featuredPrices, highlightedPrices)
//	fmt.Println(prices)
//	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
//	highlightedPrices = highlightedPrices[:3]
//	fmt.Println(highlightedPrices)
//	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
//}
