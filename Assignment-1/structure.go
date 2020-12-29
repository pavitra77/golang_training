package main

import "fmt"

type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {

	var a Address
	fmt.Println(a)

	a1 := Address{"xyz", "bombay", 689}

	fmt.Println("Address1: ", a1)

	a2 := Address{Name: "yxz", city: "delhi",
		Pincode: 401}

	fmt.Println("Address2: ", a2)

	a3 := Address{Name: "zyx"}

	fmt.Println("Address3: ", a3)
}