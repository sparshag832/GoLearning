package main

import (
	"fmt"
)

func main() {
	fmt.Println("Today We Will Learn Maps In Go Lang")

	studentGrades := make(map[string]int) // Creation Of Unordered Map
	studentGrades["Tarush"] = 12
	studentGrades["Sathak"] = 98
	studentGrades["Prateek"] = 76
	studentGrades["Sparsh"] = 99
	studentGrades["Vishal"] = 45

	fmt.Println("Marks Of Sparsh Is", studentGrades["Sparsh"])

	for index, value := range studentGrades {
		fmt.Printf("Index Is %s and value is %d\n", index, value) //Iterate Using Range
	}

	studentGrades["Sarthak"] = 54 // Change value
	fmt.Println("Marks Of Sarthak Is", studentGrades["Sarthak"])

	delete(studentGrades, "Sarthak")
	fmt.Println("Marks Of Sarthak Is", studentGrades["Sarthak"])

	//Check Whether The Record Present Or not

	grades, exists := studentGrades["Sarthak"]
	fmt.Printf("Grades Is %d\n", grades)
	fmt.Printf("IsExists %t\n", exists)

	person := map[string]int{ // Creation and initialization of Map
		"Alice":  31,
		"Bob":    12,
		"Sparsh": 34,
	}

	for keys, marks := range person {
		fmt.Printf("Key is %s and marks is %d\n", keys, marks)
	}

}
