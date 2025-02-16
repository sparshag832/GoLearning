package main

import "fmt"

func main() {
	fmt.Println("Today We Will learn For Loops")
	x := 5

	for i := 1; i <= 10; i++ {
		fmt.Println("Number Is ", i)
	}

	for {
		fmt.Println("This Is Infinite ", x)
		x--
		if x < 2 {
			fmt.Println("This is the break statement")
			break
		}
	}

	//Range Is used to give index and values of collections
	numbers := [5]int{3, 6, 2, 9, 4}
	for index, value := range numbers {
		fmt.Printf("Index Is :%d  and Value Is :%d \n", index, value)
	}

	//Range Can also be used for String and Characters
	data := "Hello World"
	for index, char := range data {
		fmt.Printf("Index Is :%d  and Value Is :%c \n", index, char)
	}
}
