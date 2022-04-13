package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	myJson := `[
					{
						"first_name": "Clark",
						"last_name": "Kent",
						"hair_color": "black",
						"has_dog": true
					},
					{
						"first_name": "Bruce",
						"last_name": "Wayne",
						"hair_color": "black",
						"has_dog": false
					}
				]`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)
	// Write JSON from a Struct
	var mySlice []Person
	var m1 = Person{"Wally", "West", "Blue", false}
	mySlice = append(mySlice, m1)
	var m2 = Person{"Peter", "Parker", "Black", false}
	mySlice = append(mySlice, m2)

	newJson, err := json.Marshal(mySlice)
	if err != nil {
		log.Println("Error marshalling json", err)
	}
	fmt.Print(string(newJson))
}
