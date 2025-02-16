package main

import "fmt"

type Person struct {
	fName string
	lName string
	age   int
}

func main() {
	var person1 Person // Declaration Of Person     1st Method

	person1.age = 21 // Initialization Of Person
	person1.fName = "Ravi"
	person1.lName = "Sharma"
	fmt.Println("This Is The First Person :", person1)

	person2 := Person{ // Declaration and Initialization    2nd Method
		fName: "Sahil",
		lName: "Khan",
		age:   23,
	}
	fmt.Println("This Is The Second Person :", person2)

	var person3 = new(Person) // 3rd Method , when we use new , it acts as a pointer
	person3.age = 22          // Initialization Of Person
	person3.fName = "Bob"
	person3.lName = "Sharma"
	fmt.Println("This Is The Third Person :", person3)
}
