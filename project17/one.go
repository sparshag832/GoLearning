package main

import (
	"encoding/json"
	"fmt"
)

// Marsheling or Encoding

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	fmt.Println("Today We are learning JSON In Go")
	person := Person{Name: "John", Age: 34, IsAdult: true}
	fmt.Println("Person Details Are: ", person)

	json.Marshal(person)
	jsonDate, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error In Marshel ", err)
	}
	fmt.Println("Preson Details In JSON Format: ", string(jsonDate))

	// UnMarsheling or Decoding

	var decodedData Person
	err = json.Unmarshal(jsonDate, &decodedData)
	if err != nil {
		fmt.Println("Error In Un Marshel ", err)
		return
	}
	fmt.Println("person data after unmarshel: ", decodedData)
}
