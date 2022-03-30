package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

// printDetails function prints the product data
func (product *Product) printDetails() {
	fmt.Printf("Product - %v (%v) is %v, priced at %.2f USD.", product.title, product.id, product.description, product.price)
}

// store function saves the product content to a file
func (product *Product) store() {
	file, _ := os.Create(product.id + ".txt")
	content := fmt.Sprintf("Product - %v is %v, priced at %.2f USD.", product.title, product.description, product.price)
	_, _ = file.WriteString(content)
	_ = file.Close()

}

func NewProduct(id string, title string, description string, price float64) *Product {
	return &Product{id, title, description, price}
}

func main() {
	createdProduct := getProduct()
	createdProduct.printDetails()
	createdProduct.store()
}

func getProduct() *Product {
	fmt.Println("Please enter the product data.")
	fmt.Println("----------------------")
	reader := bufio.NewReader(os.Stdin)
	idInput := readUserInput(reader, "Product id: ")
	titleInput := readUserInput(reader, "Product Title: ")
	descriptionInput := readUserInput(reader, "Product Description: ")
	priceInput := readUserInput(reader, "Product Price: ")
	price, _ := strconv.ParseFloat(priceInput, 64)
	product := NewProduct(idInput, titleInput, descriptionInput, price)
	return product

}

func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	return userInput
}
