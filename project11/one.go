package main

import (
	"fmt"
	"strings"
)

func main() {

	data := "Bottle,Bag,Chair,Pen,Copy"
	parts := strings.Split(data, ",") // Split function
	fmt.Println(parts)
	fmt.Printf("Type Of parts if %T\n", parts) // Array Of string

	str1 := "one two three one four five four one"
	count1 := strings.Count(str1, "one") // Count function
	fmt.Println("Count of one is", count1)
	count2 := strings.Count(str1, "four")
	fmt.Println("Count of four is", count2)

	str2 := "      hii bro     How Are U        " // Trim extra space
	trimmed := strings.TrimSpace(str2)
	fmt.Println(trimmed)

	str3 := "Sparsh"
	str4 := "Agarwal"
	result := strings.Join([]string{str3, "Kumar", str4}, " ") // Concatination
	fmt.Println("After Concatinating The Name Is :", result)

}
