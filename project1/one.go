package main

import "fmt"

func main() {
	fmt.Println("Hii")
	fmt.Println("My Name Is Sparrsh")

	var name = "Vishal"
	const age = 22

	fmt.Println("Welcome from "+name+" My Age Is", age)

	fmt.Printf("Hii My Age Is %v", age)

	fmt.Println()

	// We will make a user input

	var userName string
	var userAge int
	userName = "Ramesh"
	userAge = 12

	var user1 int
	fmt.Scan(&user1)

	fmt.Printf("User Name Is %s \n ", userName)
	fmt.Println("Hii My UserAge IS ", userAge)

	var user2 string
	fmt.Println("Enter Your UserName2")
	fmt.Scan(&user2)

	fmt.Println("My User 2 Name Is:", user2)

}
