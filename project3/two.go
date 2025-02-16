package main

import (
	"fmt"
	"project3/util"
)

func main() {
	sum := util.Sum(1, 5)
	fmt.Println("Sum IS ", sum)

	mult := util.Mult(6, 8)
	fmt.Println("Multiplication Is ", mult)

	divide, _ := util.Divide(10, 5)
	fmt.Println("Division Result Is ", divide) // if u do not want to handle the second return the use underscore operator

	divide, err := util.Divide(10, 0)
	if err != nil {
		fmt.Println("Can't Divide By Zero")
	}
	fmt.Println("Division Result Is ", divide)
}
