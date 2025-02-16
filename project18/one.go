package main

import (
	"fmt"
	"time"
)

func sayHi() {
	fmt.Println("First Line Of Hii Function")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Last Line Of Hii Function")
}

func sayHello() {
	fmt.Println("First Line Of Hello Function")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("First Last Of Hello Function")
}
func main() {
	fmt.Println("First Line Of main function")
	go sayHello()
	sayHi()
	// time.Sleep(1000 * time.Millisecond)
}
