package main

import "fmt"

type Personal struct {
	fName string
	lName string
	age   int
}

type Contact struct {
	email string
	phone string
}

type Address struct {
	house int
	area  string
	city  string
}

type Employee struct {
	E_Personal Personal
	E_Contact  Contact
	E_Address  Address
}

func main() {
	var employee1 Employee
	employee1.E_Personal = Personal{
		fName: "Sparsh",
		lName: "Agarwal",
		age:   22,
	}

	employee1.E_Contact.email = "sp2913108@gmail.com"
	employee1.E_Contact.phone = "8946385943"

	employee1.E_Address = Address{
		house: 12,
		area:  "Alamgiriganj",
		city:  "Bareilly",
	}

	fmt.Println()

	fmt.Println("Personal Details Of Employee 1 Are As Follows :")
	fmt.Printf("First Name : %s \nLast Name : %s \nAge : %d \n", employee1.E_Personal.fName, employee1.E_Personal.lName, employee1.E_Personal.age)

	fmt.Println()

	fmt.Println("Contact Details Of Employee 1 Are As Follows :")
	fmt.Printf("Email : %s \nPhone : %s \n", employee1.E_Contact.email, employee1.E_Contact.phone)

	fmt.Println()

	fmt.Println("Address Of Employee 1 Is :")
	fmt.Printf("House No. : %d \nArea: %s \nCity : %s\n", employee1.E_Address.house, employee1.E_Address.area, employee1.E_Address.city)

	fmt.Println()
}
