package main

import "fmt"

func main() {
	fmt.Println(" Statement 1 : In this lesson we  will learn defer keyword")
	defer fmt.Println(" Statement 2: In this lesson we  will learn defer keyword") // will be printed at the end of program
	fmt.Println(" Statement 3 : In this lesson we  will learn defer keyword")
	defer fmt.Println(" Statement 4 : In this lesson we  will learn defer keyword")
	fmt.Println(" Statement 5 : In this lesson we  will learn defer keyword")

	// If we have two differ statement in the same function , then it will be executed as per LIFO
}
