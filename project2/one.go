package main

import "fmt"

func main() {
	var books = [3]string{"Wings", "Sting"} // Declare and initializa an array
	var bookings [50]string                 // Declare a array

	fmt.Println("I have a book named as ", books[0])

	bookings[0] = "Nana"
	bookings[1] = "Nani"

	fmt.Println("My Friend's Name Is ", bookings[0])

	fmt.Println("Whole Array Is ", books)
	fmt.Println("First Element Of Array ", books[0])
	fmt.Println("Size of Array ", len(books))
	fmt.Printf("Type Is %T", books)

	//  Now we learn Dynamic Sized Array

	fmt.Println("Now This Is dynamic array known as Slice")

	var family []string

	family = append(family, "Mamma")
	family = append(family, "Papa")
	family = append(family, "sis")

	fmt.Println("Whole Array Is ", family)
	fmt.Println("First Element Of Array ", family[0])
	fmt.Printf("Type Is %T", family)

	// for {                                        Infinite loop
	// 	fmt.Println("HII")
	// }

}
