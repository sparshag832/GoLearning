package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Data Conversion

	var num int = 42
	fmt.Printf("Type Of num is %T\n", num)

	var num1 float64 = float64(num)

	fmt.Println("Sum Is", num1+12.4)

	str := strconv.Itoa(num) // anything to string
	fmt.Printf("DataType Of str is %T\n", str)

	number_string := "1234"
	number_int, _ := strconv.Atoi(number_string) // Convert numberString to integer number
	fmt.Println("Value Of number_int is ", number_int)
	fmt.Printf("Type Of number is %T\n", number_int)

	num_string := "3.14"
	number_Float, _ := strconv.ParseFloat(num_string, 64) //String To Float
	fmt.Println("Value Of number_int is ", number_Float)
	fmt.Printf("Type Of number is %T\n", number_Float)

}
