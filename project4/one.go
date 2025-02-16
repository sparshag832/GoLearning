package main

import "fmt"

func main() {
	fmt.Println("Today We Will learn If else and switch")

	x := 8
	if x > 6 {
		fmt.Println("This is the example of If statement")
	} else {
		fmt.Println("This is else statement")
	}

	switch x {
	case 1, 8:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Unknown")
	}

}
