package main

import "fmt"

func main() {

	//Creating Channel
	var mychannel chan int
	fmt.Println("Value Of the channel: ", mychannel)
	fmt.Printf("Type Of Channel is: %T", mychannel)

	println()

	// Creating a channel using make() function
	mychannel1 := make(chan int)
	fmt.Println("\nValue of the channel1: ", mychannel1)
	fmt.Printf("Type of the channel1: %T ", mychannel1)

}

func myfunc(ch chan int) {
	fmt.Println(235 + <-ch)
}
