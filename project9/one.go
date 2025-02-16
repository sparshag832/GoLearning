package main

import "fmt"

func modifyValueByReference(ptr2 *int) {
	*ptr2 = *ptr2 + 5
}

func main() {
	num := 2

	var ptr *int

	// var ptr1 *int

	ptr = &num

	fmt.Println("Address Of num that is stored in ptr is", ptr)
	fmt.Println("Value At Address stored in ptr is ", *ptr)

	// fmt.Println(ptr1) // If not assigned to anyone then nil will be given as result

	modifyValueByReference(ptr)

	fmt.Println("After Modification Value At Address stored in ptr is ", num)

}
